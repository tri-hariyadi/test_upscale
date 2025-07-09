package main

import (
	"errors"
	"net/http"
	"todo-app/common/config"
	"todo-app/common/injector"
	"todo-app/common/logger"
	httpserver "todo-app/common/server"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	channel := make(chan *http.Server)
	channel2 := make(chan func())
	defer close(channel)
	defer close(channel2)

	go func() {
		server, cleanup := injector.InitializedServer()
		channel <- server
		channel2 <- cleanup
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatalf("Server error: %v", err)
		}
	}()

	srv := <- channel
	cleanup := <- channel2
	httpserver.GraceFullShutDown(srv, cleanup)
}
