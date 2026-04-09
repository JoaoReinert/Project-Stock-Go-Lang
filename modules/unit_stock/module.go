package unitStock

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type moduleUnitStock struct {
	useCase domain.UnitStockUseCase
	path    string
	name    string
}

func NewUnitStockModule(
	useCase domain.UnitStockUseCase,
) modules.Module {
	return moduleUnitStock{
		useCase: useCase,
		name:    "Unit stock module",
		path:    "/unitStock",
	}
}

func (m moduleUnitStock) Name() string {
	return m.name
}

func (m moduleUnitStock) Path() string {
	return m.path
}

func (m moduleUnitStock) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerUnitStock,
			Path:    "/register",
			Label:   "Register unit stock in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.updateUnitStock,
			Path:    "/update",
			Label:   "Update unit stock in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.deleteUnitStock,
			Path:    "/delete/{id}",
			Label:   "Delete unit stock in database",
			Methods: []string{http.MethodDelete},
		},
		{
			Handler: m.detailsUnitStock,
			Path:    "/details/{id}",
			Label:   "Details unit stock in database",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleUnitStock) registerUnitStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var unitStock entities.UnitStock
	err = json.Unmarshal(body, &unitStock)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [RegisterUnitStock]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleUnitStock) updateUnitStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var unitStock entities.UnitStock
	err = json.Unmarshal(body, &unitStock)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.UpdateUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [UpdateUnitStock]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleUnitStock) deleteUnitStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unitStock := entities.UnitStock{ID: int64(id)}

	err = m.useCase.DeleteUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [DeleteUnitStock]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleUnitStock) detailsUnitStock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unitStock := entities.UnitStock{
		ID: int64(id),
	}

	details, err := m.useCase.DetailsUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [DetailsUnitStock]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(details)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
