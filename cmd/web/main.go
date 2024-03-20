package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "Server address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	tc, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        slog.New(slog.NewTextHandler(os.Stdout, nil)),
		templateCache: tc,
	}

	srv := http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("started server", "addr", *addr)

	err = srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}
