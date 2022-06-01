package rproxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nyan2d/redirector/config"
)

type BaseHandler struct {
	proxy *RProxy
}

func (bh *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	if proxy, ok := bh.proxy.proxies[host]; ok {
		proxy.ServeHTTP(w, r)
		return
	}

	if targetHost, ok := bh.proxy.hosts[host]; ok {
		remoteHost, err := url.Parse(targetHost)
		if err != nil {
			log.Println("parsing target host:", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteHost)
		bh.proxy.proxies[host] = proxy
		proxy.ServeHTTP(w, r)
		return
	}

	errorMessage := "403: Forbidden"
	w.Write([]byte(errorMessage))
}

type RProxy struct {
	hosts   map[string]string
	proxies map[string]*httputil.ReverseProxy
}

func NewRProxy(config *config.Config) *RProxy {
	return &RProxy{
		hosts:   config.HostsMap(),
		proxies: make(map[string]*httputil.ReverseProxy),
	}
}

func (rp *RProxy) Listen(addr string) error {
	handler := &BaseHandler{
		proxy: rp,
	}

	httpserver := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return httpserver.ListenAndServe()
}
