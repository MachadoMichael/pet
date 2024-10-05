package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const port = ":4000"

func Init() (*http.Server, error) {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	// add handlers and middlewares
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	})

	srv := &http.Server{Addr: port, Handler: r}

	return srv, nil
}

func Start() {
	server, err := Init()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Starting server on port: %s  ", port)
	server.ListenAndServe()
}
