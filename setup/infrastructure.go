package setup

import (
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"Desafio_Go_Lang/modules/authentication"
	"Desafio_Go_Lang/modules/category"
	"Desafio_Go_Lang/modules/settings"
	"Desafio_Go_Lang/modules/subCategory"
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

	categoryRepository := category.NewCategoryRepository(database)

	categoryUseCase := category.NewCategoryUseCase(categoryRepository, cfg)

	categoryModule := category.NewCategoryModule(categoryUseCase)

	subCategoryRepository := subCategory.NewSubCategoryRepository(database)

	subCategoryUseCase := subCategory.NewSubCategoryUseCase(subCategoryRepository, cfg)

	subCategoryModule := subCategory.NewSubCategoryModule(subCategoryUseCase)

	applicationModules := []modules.Module{
		categoryModule,
		subCategoryModule,
	}

	routerBase := authenticationModule.Setup(r)
	for _, am := range applicationModules {
		moduleSubRouter := routerBase.PathPrefix(am.Path()).Subrouter()
		_ = am.Setup(moduleSubRouter)
	}
}
