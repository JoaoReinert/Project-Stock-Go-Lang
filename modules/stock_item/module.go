package stock_item

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/domain/util"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type moduleStockItem struct {
	useCase domain.StockItemUseCase
	path    string
	name    string
}

func NewStockItemModule(
	useCase domain.StockItemUseCase,
) modules.Module {
	return moduleStockItem{
		useCase: useCase,
		name:    "Stock item module",
		path:    "/stockItem",
	}
}

func (m moduleStockItem) Name() string {
	return m.name
}

func (m moduleStockItem) Path() string {
	return m.path
}

func (m moduleStockItem) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.addStockItem,
			Path:    "/add",
			Label:   "Add stock item in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.removeStockItem,
			Path:    "/remove",
			Label:   "Remove stock items in database",
			Methods: []string{http.MethodPost},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleStockItem) addStockItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := util.GetUser(r)
	if err != nil {
		log.Printf("Error in [GetUser]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Erro in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stockItem entities.StockItem
	err = json.Unmarshal(body, &stockItem)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.AddStockItem(ctx, stockItem, user.ID)
	if err != nil {
		log.Printf("Error in [RegisterStockItem]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleStockItem) removeStockItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := util.GetUser(r)
	if err != nil {
		log.Printf("Error in [GetUser]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Erro in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stockItem entities.StockItem
	err = json.Unmarshal(body, &stockItem)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RemoveStockItem(ctx, stockItem, user.ID)
	if err != nil {
		log.Printf("Error in [RemoveStockItem]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
