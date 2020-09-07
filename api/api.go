package api

import (
	"english-verbs/verbs/app"
	http2 "english-verbs/verbs/infraestructure/http"
	"english-verbs/verbs/infraestructure/persistence"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK from Container"))
}

func Start(port string) {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/heartbeat", healthCheck)

	accountController := accountController()

	r.Mount("/verbs", accountController.Routes())

	server := newServer(port, r)

	server.Start()

}

func AllowOriginFunc(r *http.Request, origin string) bool {

	return true
}

func accountController() http2.VerbHandler {

	repo := persistence.VerbRepository{}

	servicio := app.VerbService{Repo: &repo}

	return http2.VerbHandler{Service: servicio}

}
