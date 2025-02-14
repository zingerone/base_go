package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zingerone/base_go/internal/server/route"
	"log"
	"net/http"
)

// RunAPI runs the API server

func RunAPI(ctx context.Context) {

	r := gin.Default()
	route.SetupRouter(r)

	// Graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		select {
		case <-ctx.Done():
			log.Println("Shutting down server...")
		}

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
