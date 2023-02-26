package chi

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

func Initialize() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", ListUsers)
		r.Post("/", CreateUser)

		r.Route("/{email}", func(r chi.Router) {
			r.Get("/", GetUser)
			r.Delete("/", DeleteUser)
			r.Patch("/", UpdateUser)
		})
	})

	return r
}
