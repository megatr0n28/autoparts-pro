package main

import (
	"log"

	"net/http"

	"os"

	"os/signal"

	"syscall"

	"context"

	"time"

	apihttp "github.com/megatr0n28/autoparts-pro/backend/internal/interfaces/http"
)

func main() {

	router := apihttp.NewRouter()

	server := &http.Server{

		Addr: ":8080",

		Handler: router,
	}

	go func() {

		log.Println(
			"AutoParts Pro API started on :8080",
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatal(err)

		}

	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println(
		"Shutting down server...",
	)

	ctx, cancel :=
		context.WithTimeout(
			context.Background(),
			10*time.Second,
		)

	defer cancel()

	if err :=
		server.Shutdown(ctx); err != nil {

		log.Fatal(
			"Server forced shutdown",
		)

	}

	log.Println(
		"Server stopped",
	)

}
