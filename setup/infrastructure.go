package setup

import (
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"Desafio_Go_Lang/modules/authentication"
	"Desafio_Go_Lang/modules/settings"
	"log/slog"

	"github.com/gorilla/mux"
)

func SetupModules(r *mux.Router, cfg entities.Config) {
	slog.Info("Setup modules")

	database := settings.NewSettingsRepository(cfg)
	database.Connection()

	authenticationRepository := authentication.NewAuthenticationRepository(database)

	authenticationUseCase := authentication.NewAuthenticationUseCase(authenticationRepository, cfg.Paseto.PasetoSecurityKey, cfg.Paseto.UserPassSaltSecret)

	authenticationModule := authentication.NewAuthenticationModule(cfg, authenticationUseCase)

	applicationModules := []modules.Module{}

	routerBase := authenticationModule.Setup(r)
	for _, am := range applicationModules {
		moduleSubRouter := routerBase.PathPrefix(am.Path()).Subrouter()
		_ = am.Setup(moduleSubRouter)
	}
}
