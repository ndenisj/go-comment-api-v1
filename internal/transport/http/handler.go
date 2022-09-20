package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the routes in our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up our Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
