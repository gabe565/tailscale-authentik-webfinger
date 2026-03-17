package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"gabe565.com/tailscale-authentik-webfinger/internal/config"
	"gabe565.com/tailscale-authentik-webfinger/internal/webfinger"
	"gabe565.com/utils/bytefmt"
	"gabe565.com/utils/versionx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var version = "beta"

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}

	vers := versionx.New(version)

	slog.Info("Tailscale Authentik Webfinger",
		"version", vers.Version,
		"commit", vers.Commit.Short(),
	)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	if conf.RealIPHeader {
		r.Use(middleware.RealIP)
	}
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.GetHead)

	r.Get("/.well-known/webfinger", webfinger.Handler(conf))

	server := &http.Server{
		Addr:           conf.ListenAddress,
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: bytefmt.MiB,
	}
	slog.Info("Listening on " + server.Addr)
	return server.ListenAndServe()
}
