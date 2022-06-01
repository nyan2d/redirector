package main

import (
	"log"

	"github.com/nyan2d/redirector/config"
	"github.com/nyan2d/redirector/rproxy"
)

func main() {
	cfg, err := config.ReadFromFile("config.yaml")
	if err != nil {
		log.Fatal("reading config:", err)
	}

	proxy := rproxy.NewRProxy(cfg)
	err = proxy.Listen(cfg.Address)
	log.Println(err)
}
