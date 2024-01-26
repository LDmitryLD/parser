package router

import (
	"net/http"
	"projects/LDmitryLD/parser/app/internal/modules"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()
	setDefaultRoutes(r)

	r.Post("/parser/search", controllers.Vacancy.Search)
	r.Get("/parser/{id}", controllers.Vacancy.Get)
	r.Get("/parser/list", controllers.Vacancy.List)
	r.Delete("/parser/delete/{id}", controllers.Vacancy.Delete)

	return r
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})
}
