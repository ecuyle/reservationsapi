package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
    "github.com/ecuyle/reservationsapi/routes"
)

func Routes() *chi.Mux {
    router := chi.NewRouter();
    router.Use(
        render.SetContentType(render.ContentTypeJSON),
        middleware.DefaultCompress,
        middleware.RedirectSlashes,
        middleware.Recoverer,
    )
    //took out middleware.Logger

    router.Route("/", func(r chi.Router) {
        r.Mount("/api/reserve", routes.Routes())
    })

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

    log.Fatal(http.ListenAndServe(":3001", router))
}

