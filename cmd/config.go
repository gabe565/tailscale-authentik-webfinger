package cmd

import "errors"

type Config struct {
	ListenAddress string `name:"listen-address" description:"Address to listen on"                  default:":3000"`
	Domain        string `name:"domain"         description:"Authentik domain"`
	AppName       string `name:"app-name"       description:"Authentik app name"                    default:"tailscale"`
	RealIPHeader  bool   `name:"real-ip-header" description:"Get client IP from the Real-IP header" default:"true"`
}

var ErrDomain = errors.New("domain is required")

func (c Config) Validate() error {
	if c.Domain == "" {
		return ErrDomain
	}
	return nil
}
