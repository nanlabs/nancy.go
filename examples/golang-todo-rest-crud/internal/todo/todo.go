// todo package include application logic related with the todo feature.
package todo

import (
	"net/http"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/platform/mongo"
	"todo-api-golang/internal/todo/note"
	"todo-api-golang/internal/trace"
	"todo-api-golang/pkg/health"
	"todo-api-golang/pkg/logs"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// APIServer is the handles the setup of a todo api.
type APIServer struct {
	NoteHandler note.Handler
	Router      *mux.Router
	Logs        *logs.Logs
}

// NewApi creates the default configuration for the http server and set up routing.
func NewApi(config *config.Config, mongoClient mongo.ClientHelper, logs *logs.Logs) (*APIServer, error) {

	noteRepository := note.NewRepository(mongoClient, config)
	noteService := note.NewService(noteRepository)
	validator := validator.New()
	if err := validator.RegisterValidation("enum", note.ValidateEnum); err != nil {
		logs.Logger.Error("Failed registering validators for handlers", zap.String("detail", err.Error()))
		return nil, err
	}
	noteHandler := note.NewHandler(noteService, validator)

	router := setupRoutes(noteHandler, logs, config)

	return &APIServer{
		NoteHandler: noteHandler,
		Router:      router,
		Logs:        logs,
	}, nil
}

// setupRoutes create the routes for the todo api.
func setupRoutes(noteHandler note.Handler, logs *logs.Logs, config *config.Config) *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.Use(trace.ContextIDMiddleware(logs))
	base.Use(LogMiddleware(logs))

	base.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.GetAll).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.Create).Methods(http.MethodPost)
	base.HandleFunc("/notes/{noteId}", noteHandler.GetById).Methods(http.MethodGet)
	base.HandleFunc("/notes/{noteId}", noteHandler.Update).Methods(http.MethodPatch)

	return base
}
