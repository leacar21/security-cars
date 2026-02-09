package usecase

import "github.com/example/security-cars/backend/internal/domain"

type GreetingService struct {}

func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

func (service *GreetingService) GetGreeting() string {
	greeting := domain.Greeting{Message: "Hello from the Go backend"}
	return greeting.Message
}
