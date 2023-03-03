package http

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/monirz/goserve"
	"github.com/monirz/goserve/config"
	"github.com/monirz/goserve/postgres"

	"github.com/go-chi/chi"
)

type Server struct {
	db     *sql.DB
	router *chi.Mux
	Config *config.Config

	UserService goserve.UserService
}

func NewServer() *Server {

	// var err error
	s := &Server{}
	s.Config = config.NewConfig()

	s.router = chi.NewRouter()

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	s.router.Route("/api/v1", func(r chi.Router) {

	})

	return s

}

func (s *Server) Run() {

	s.UserService = postgres.NewUserService(s.db)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	addr := fmt.Sprintf(":%s", s.Config.Port)

	srv := &http.Server{
		Addr: addr,
		// ReadTimeout:  60 * time.Second,
		// WriteTimeout: 60 * time.Second,
		Handler: s.router,
	}

	go func() {
		log.Println("Staring server with address ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Failed to start http server on :", err)
			os.Exit(-1)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Http server couldn't shutdown gracefully")
	}

	log.Println("shutting down")

}
