package routes

import (
	"jwt/handlers"

	"github.com/go-chi/chi"
)

func Routes(router *chi.Mux) {
	// group other routes with /api
	router.Route("/api", func(r chi.Router) {
		UserRoutes(r)
		LoginRoutes(r)
	})
}

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.UserListHandler)
		r.Get("/{id}", handlers.UserGetByIDHandler)
		r.Post("/", handlers.UserInsertHandler)
		r.Put("/{id}", handlers.UpdateUserHandler)
		r.Delete("/{id}", handlers.DeleteUserHandler)
	})
}
func LoginRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.LoginHandler)
		// r.Get("/{id}", handlers.UserGetByIDHandler)
		// r.Post("/", handlers.UserInsertHandler)
		// r.Put("/{id}", handlers.UpdateUserHandler)
		// r.Delete("/{id}", handlers.DeleteUserHandler)
	})
}
