package cmd

import "errors"

type Config struct {
	ListenAddress string `name:"listen-address" description:"Address to listen on"                  default:":3000"`
	AuthentikHost string `name:"ak-host"        description:"Authentik domain"`
	AuthentikApp  string `name:"ak-app-name"    description:"Authentik app name"                    default:"tailscale"`
	RealIPHeader  bool   `name:"real-ip-header" description:"Get client IP from the Real-IP header" default:"true"`
}

var ErrDomain = errors.New("ak-host is required")

func (c Config) Validate() error {
	if c.AuthentikHost == "" {
		return ErrDomain
	}
	return nil
}
