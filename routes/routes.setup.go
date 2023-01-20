package routes

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	// "github.com/parinyapt/PT-Friendship_Backend/handler"
)

func Setup() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	router := gin.Default()
	//config
	configCors(router)
	configRateLimit(router)
	s := configApi(router, os.Getenv("PORT"))

	//setup all api route
	configApiRoutes(router)

	// Gracefully shutdown the server
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Error]->ListenAndServe Fail : %s", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 60 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("[Error]->Server forced to shutdown fail: ", err)
	}

	log.Println("Server shutting down completely")
}
