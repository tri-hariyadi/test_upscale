package server

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"strconv"
	"todo-app/common/config"
	"todo-app/common/logger"
	"todo-app/middleware"
)

func NewCleanupFunc(db *pgxpool.Pool) func() {
	return func() {
		logger.Logger.Info("Closing a database connection...")
		db.Close()
		logger.Logger.Info("Database connection closed.")
	}
}

func NewServer(handler *mux.Router, db *pgxpool.Pool) (*http.Server, func()) {
	logger.Logger.Info("\033[34mInitializing HTTP Server with Port : "+strconv.Itoa(config.AppConfig.App.Port)+"\033[0m")
	apiUrl := config.AppConfig.App.Host + ":" + strconv.Itoa(config.AppConfig.App.Port)
	storageUrl := config.AppConfig.Storage.Host + ":" +strconv.Itoa(config.AppConfig.Storage.Port)
	logger.Logger.Infof("\033[34mAPI url: http://%s\033[0m", apiUrl)
	logger.Logger.Infof("\033[34mStorage url: http://%s\033[0m", storageUrl)

	go func() {
		directory := flag.String("d", ".", "the directory of static file to host")
		flag.Parse()

		http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Incoming request:", r.Method, r.URL.Path, r.Header.Get("User-Agent"))
			http.FileServer(http.Dir(*directory)).ServeHTTP(w, r)
		})
		if err := http.ListenAndServe(storageUrl, nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatalf("Server error: %v", err)
		}
	}()

	return &http.Server{
		Addr:    apiUrl,
		Handler: middleware.CORSMiddleware(handler),
	}, NewCleanupFunc(db)
}