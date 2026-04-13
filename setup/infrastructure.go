package setup

import (
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"Desafio_Go_Lang/modules/authentication"
	"Desafio_Go_Lang/modules/category"
	"Desafio_Go_Lang/modules/equipment"
	"Desafio_Go_Lang/modules/settings"
	"Desafio_Go_Lang/modules/stock_item"
	"Desafio_Go_Lang/modules/sub_category"
	unitStock "Desafio_Go_Lang/modules/unit_stock"
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

	subCategoryRepository := sub_category.NewSubCategoryRepository(database)

	subCategoryUseCase := sub_category.NewSubCategoryUseCase(subCategoryRepository, cfg)

	subCategoryModule := sub_category.NewSubCategoryModule(subCategoryUseCase)

	equipmentRepository := equipment.NewEquipmentRepository(database)

	equipmentUseCase := equipment.NewEquipmentUseCase(equipmentRepository, cfg)

	equipmentModule := equipment.NewEquipmentModule(equipmentUseCase)

	unitStockRepository := unitStock.NewUnitStockRepository(database)

	unitStockUseCase := unitStock.NewUnitStockUseCase(unitStockRepository, cfg)

	unitStockModule := unitStock.NewUnitStockModule(unitStockUseCase)

	stockItemRepository := stock_item.NewStockItemRepository(database)

	stockItemUseCase := stock_item.NewStockItemUseCase(stockItemRepository, cfg)

	stockItemModule := stock_item.NewStockItemModule(stockItemUseCase)

	applicationModules := []modules.Module{
		categoryModule,
		subCategoryModule,
		equipmentModule,
		unitStockModule,
		stockItemModule,
	}

	routerBase := authenticationModule.Setup(r)
	for _, am := range applicationModules {
		moduleSubRouter := routerBase.PathPrefix(am.Path()).Subrouter()
		_ = am.Setup(moduleSubRouter)
	}
}
