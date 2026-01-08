// Package server provides functionalities related to the web server.
package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nelsonmarro/go-stripe/config"
)

type Server struct {
	Config   *config.Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DB       *sql.DB
}

func NewServer(cfg *config.Config, infoLog, errorLog *log.Logger, db *sql.DB) *Server {
	return &Server{
		Config:   cfg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		DB:       db,
	}
}

func (s *Server) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.Config.Port),
		Handler:           s.getRoutes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	s.InfoLog.Println("Starting server on port", s.Config.Port)

	return srv.ListenAndServe()
}
