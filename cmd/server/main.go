package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Quak1/chuy-gbf/internal/database"
	"github.com/Quak1/chuy-gbf/internal/routes"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Failed to start DB: %v\n", err)
		return
	}
	defer db.Close()

	r := routes.SetupRouter(db)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		log.Printf("Shutting down server: %s\n", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	log.Printf("Starting server on port %s\n", srv.Addr)

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server failed: %v\n", err)
		return
	}

	err = <-shutdownError
	if err != nil {
		log.Printf("Error shutting server down: %v\n", err)
	}

	log.Println("Server stopped")
}
