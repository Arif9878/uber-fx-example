package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

// Define a service interface
type GreeterService interface {
	Greet() string
}

// Implement a service
type EnglishGreeter struct{}

func (e *EnglishGreeter) Greet() string {
	return "Hello, World!"
}

// Define a function to provide the GreeterService
func ProvideGreeter() GreeterService {
	return &EnglishGreeter{}
}

func main() {
	// Create a new FX application
	app := fx.New(
		// Provide the GreeterService using the ProvideGreeter function
		fx.Provide(ProvideGreeter),
		// Define an "invoke" function to consume the GreeterService
		fx.Invoke(func(greeter GreeterService) {
			message := greeter.Greet()
			fmt.Println(message)
		}),
	)

	// Run the application
	if err := app.Start(context.Background()); err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}

	// Wait for the application to be shut down (e.g., using Ctrl+C)
	<-app.Done()
}
