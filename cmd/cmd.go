package cmd

import (
	"context"
	"net/http"
	"time"

	"gabe565.com/utils/bytefmt"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/must"
	"github.com/USA-RedDragon/configulator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authentik-webfinger",
		Short: "Provides an Authentik webfinger endpoint",
		RunE:  run,

		SilenceErrors: true,
	}

	cmd.SetContext(
		configulator.New[Config]().
			WithPFlags(cmd.Flags(), nil).
			WithEnvironmentVariables(&configulator.EnvironmentVariableOptions{
				Prefix:    "",
				Separator: "__",
			}).
			WithContext(context.Background()),
	)

	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	config, err := must.Must2(configulator.FromContext[Config](cmd.Context())).Load()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	if config.RealIPHeader {
		r.Use(middleware.RealIP)
	}
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.GetHead)

	r.Get("/.well-known/webfinger", handler(config))

	server := &http.Server{
		Addr:           config.ListenAddress,
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: bytefmt.MiB,
	}
	return server.ListenAndServe()
}
