package routes

import (
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/ecuyle/reservationsapi/db"
)

var conn = db.PGConnect()

type Reservation struct {
    BookingsToday int `json:"bookings_today"`
}

func Routes() *chi.Mux {
    router := chi.NewRouter()
    router.Get("/load/{id}", GetBookingsTodayForRestaurant)
    return router
}

func GetBookingsTodayForRestaurant(res http.ResponseWriter, req *http.Request) {
    id := chi.URLParam(req, "id")
    var reservation Reservation
    req.Header.Set("Connection", "close")
    queryString := "SELECT bookings_today from restaurants WHERE id=$1"
    row := conn.QueryRow(queryString, id)
    err := row.Scan(&reservation.BookingsToday)
    req.Close = true
    if err != nil {
        render.JSON(res, req, ErrInvalidRequest(err))
        return
    }
    render.JSON(res, req, reservation)
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
