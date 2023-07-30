package web

import "github.com/go-chi/chi/v5"

type Route struct {
	chi.Router
}

func NewRoute() *Route {
	r := chi.NewRouter()
	return &Route{r}
}
