package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/common/logger"
)

func GraceFullShutDown(srv *http.Server, cleanup func()) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error("Server shutdown failed:", err)
	} else {
		logger.Logger.Info("Shutting down gracefully")
	}

	cleanup()
}
