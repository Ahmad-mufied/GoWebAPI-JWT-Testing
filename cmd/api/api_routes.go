package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	// curl http://localhost:8090/auth -X POST -H "Content-Type: application/json" -d '{"email":"admin@example.com","password":"secret"}'
	mux.Post("/refresh-token", app.refresh)

	//// test handler
	//mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
	//	var payload = struct {
	//		Message string `json:"message"`
	//	}{
	//		Message: "hello, world",
	//	}
	//
	//	_ = app.writeJSON(w, http.StatusOK, payload)
	//
	//	// curl http://localhost:8090/test
	//})

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware
		mux.Use(app.authRequired)

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
