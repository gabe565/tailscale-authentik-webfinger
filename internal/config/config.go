package config

import "github.com/caarlos0/env/v11"

//go:generate go tool envdoc -types Config -output ../../envs.md

type Config struct {
	// Address to listen on
	ListenAddress string `env:"LISTEN_ADDRESS" envDefault:":3000"`
	// Authentik domain
	AuthentikHost string `env:"AK_HOST,required"`
	// Authentik app name
	AuthentikApp string `env:"AK_APP-NAME" envDefault:"tailscale"`
	// Get client IP from the Real-IP header
	RealIPHeader bool `env:"REAL_IP_HEADER" envDefault:"true"`
}

func Load() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
