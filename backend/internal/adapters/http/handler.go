package httpadapter

import (
	"encoding/json"
	"net/http"

	"github.com/example/security-cars/backend/internal/application/ports/inbound"
)

type Handler struct {
	greetingService inbound.GreetingUseCase
}

type greetingResponse struct {
	Message string `json:"message"`
}

func NewHandler(greetingService inbound.GreetingUseCase) *Handler {
	return &Handler{greetingService: greetingService}
}

func (handler *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.health)
	return mux
}

func (handler *Handler) health(writer http.ResponseWriter, request *http.Request) {
	response := greetingResponse{Message: handler.greetingService.GetGreeting()}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, "failed to encode response", http.StatusInternalServerError)
	}
}
