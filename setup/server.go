package setup

import (
	"Desafio_Go_Lang/entities"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupServer(router *mux.Router, configs entities.Config) (*http.Server, error) {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", configs.Server.Port),
		Handler: handlers.CORS(
			handlers.AllowedOrigins(configs.Server.AllowedOrigins),
			handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept"}),
			handlers.AllowedMethods([]string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}),
			handlers.AllowCredentials(),
		)(router),
		ReadTimeout:       time.Second * time.Duration(configs.Server.ReadTimeout),
		IdleTimeout:       time.Second * time.Duration(configs.Server.IdleTimeout),
		WriteTimeout:      time.Minute * time.Duration(configs.Server.WriteTimeout),
		ReadHeaderTimeout: time.Second * 15,
	}
	return server, nil
}
