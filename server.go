package main

import (
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
    "github.com/ecuyle/reservationsapi/routes"
)

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
    if strings.ContainsAny(path, "{}*") {
        panic("FileServer does not permit URL parameters.")
    }

    fs := http.StripPrefix(path, http.FileServer(root))

    if path != "/" && path[len(path)-1] != '/' {
        r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
        path += "/"
    }
    path += "*"

    r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fs.ServeHTTP(w, r)
    }))
}

func Routes() *chi.Mux {
    router := chi.NewRouter();
    router.Use(
        render.SetContentType(render.ContentTypeJSON),
        middleware.DefaultCompress,
        middleware.RedirectSlashes,
        middleware.Recoverer,
        middleware.Logger,
    )
    //took out middleware.Logger

    router.Route("/api", func(r chi.Router) {
        r.Mount("/reserve", routes.Routes())
    })

    FileServer(router, "/", http.Dir("./static"))

    return router
}

func main() {
    router := Routes()

    walk := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
        log.Printf("%s %s\n", method, route)
        return nil
    }

    if err := chi.Walk(router, walk); err != nil {
        log.Panicf("Logging err: %s\n", err.Error())
    }

    http.DefaultClient.Timeout = time.Minute * 10
    log.Fatal(http.ListenAndServe(":3000", router))
}

