// Package main implements the web server for the application.
package main

import (
	"log"
	"os"

	"github.com/nelsonmarro/go-stripe/config"
	"github.com/nelsonmarro/go-stripe/internal/driver"
	"github.com/nelsonmarro/go-stripe/internal/web/server"
)

func main() {
	cfg := config.LoadConfigOnce()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.ConnectSQL(cfg.DB.DSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			errorLog.Printf("error conneting to the database: %v", err)
		}
	}()

	s := server.NewServer(cfg, infoLog, errorLog, conn)

	err = s.Serve()
	if err != nil {
		s.ErrorLog.Println(err)
		panic(err)
	}
}
