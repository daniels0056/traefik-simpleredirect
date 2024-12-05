package traefik_simpleredirect

import (
	"context"
	"log"
	"net/http"
)

type Config struct {
	RedirectTo   string `json:"redirectTo"`
	RedirectCode int    `json:"redirectCode"`
}

func CreateConfig() *Config {
	return &Config{}
}

type SimpleRedirect struct {
	redirectTo   string
	redirectCode int
	next         http.Handler
	name         string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	log.Printf("SimpleRedirect Configs: [RedirectCode=%v, RedirectTo='%v']", config.RedirectCode, config.RedirectTo)

	return &SimpleRedirect{
		redirectTo:   config.RedirectTo,
		redirectCode: config.RedirectCode,
		next:         next,
		name:         name,
	}, nil
}

func (r *SimpleRedirect) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, r.redirectTo, r.redirectCode)
}
