// main package include the entry point of the todo api application
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/platform/mongo"
	"todo-api-golang/internal/ratelimit"
	"todo-api-golang/internal/todo"
	"todo-api-golang/pkg/logs"

	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"go.uber.org/zap"
)

// main is the entry point of the todo rest api.
func main() {
	logs, err := logs.New()
	if err != nil {
		log.Fatalf("Error initializing zap: %s\n", err)
	}

	config, err := config.LoadConfig("./../..")
	if err != nil {
		logs.Logger.Fatal("Cannot load config", zap.String("detail", err.Error()))
	}

	startHTTPServer(config, logs)
}

// startHTTP server initialize the server http.
func startHTTPServer(config *config.Config, logs *logs.Logs) {

	mongoClient, err := mongo.NewDbClient(config, logs)
	if err != nil {
		logs.Logger.Fatal("Error starting mongo client", zap.String("detail", err.Error()))
	}

	todoApi, err := todo.NewApi(config, mongoClient, logs)
	if err != nil {
		logs.Logger.Fatal("Error starting ToDo API", zap.String("detail", err.Error()))
	}

	// swagger
	todoApi.Router.PathPrefix("/swagger/").Handler(http.StripPrefix("/api/v1/swagger/", http.FileServer(http.Dir("./third_party/swagger-ui-4.11.1"))))

	// CORS
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	// // Rate limit
	ratel := ratelimit.LimitHandler(todoApi.Router, logs, config)

	// create a new server
	serverAddress := fmt.Sprintf("%s:%s", config.HTTPServerHost, config.HTTPServerPort)
	server := http.Server{
		Addr:         serverAddress,     // configure the bind address
		Handler:      cors(ratel),       // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logs.Logger.Info("Starting ToDo API server", zap.String("address", serverAddress))

		if err := server.ListenAndServe(); err != nil {
			logs.Logger.Fatal("Error starting ToDo API server", zap.String("detail", err.Error()))
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// block until a signal is received.
	sig := <-ch
	logs.Logger.Info("Shutdown signal received", zap.String("detail", sig.String()))

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		logs.Logger.Fatal("Error doing the shutdown", zap.String("detail", err.Error()))
	}

	logs.Logger.Info("Server shutdown completed")
}
