package route_sphere

import (
	"sync"
)

type Server struct{}

// Start - Starts the server
func (s Server) Start(waitGroup *sync.WaitGroup, environment *Environment) {
	defer waitGroup.Done()
	println("Starting server")
}
