package proxy

import (
	"api-gateway/internal/models"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Gateway struct {
	config *models.Config
	server *http.Server
}

func NewGateway(cfg *models.Config) *Gateway {
	return &Gateway{
		config: cfg,
	}
}

func (g *Gateway) Start() error {
	mux := http.NewServeMux()
	for _, service := range g.config.Services {
		proxy, err := g.createProxy(service)
		if err != nil {
			return fmt.Errorf("failed to create proxy for %s:%v", service.Name, err)
		}
		for _, path := range service.Paths {
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				proxy.ServeHTTP(w, r)
			})
		}
	}

	g.server = &http.Server{
		Addr:    g.config.Address, //API-Gateway listen port - receives external requests and routes to internal services
		Handler: mux,
	}

	return g.server.ListenAndServe()
}

func (g *Gateway) createProxy(service models.Service) (*httputil.ReverseProxy, error) {
	target, err := url.Parse("http://" + service.Name + ":" + service.Port)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(target) //Creates a reverse proxy that forwards requests to the targer service
	// 													 /login -> http://auth-service:50051/login
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Error proxying to %s: %v", service.Name, err)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}

	return proxy, nil
}
