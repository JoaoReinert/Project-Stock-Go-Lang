package equipment

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type moduleEquipment struct {
	useCase domain.EquipmentUseCase
	path    string
	name    string
}

func NewEquipmentModule(
	useCase domain.EquipmentUseCase,
) modules.Module {
	return moduleEquipment{
		useCase: useCase,
		name:    "Equipment module",
		path:    "/equipment",
	}
}

func (m moduleEquipment) Name() string {
	return m.name
}

func (m moduleEquipment) Path() string {
	return m.path
}

func (m moduleEquipment) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerEquipment,
			Path:    "/register",
			Label:   "Register equipment in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.registerEquipment,
			Path:    "/update",
			Label:   "Update equipment in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.deleteEquipment,
			Path:    "/delete/{id}",
			Label:   "Delete equipment in database",
			Methods: []string{http.MethodDelete},
		},
		{
			Handler: m.detailsEquipment,
			Path:    "/details/{id}",
			Label:   "Get details equipment in database",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleEquipment) registerEquipment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var equipment entities.Equipment
	err = json.Unmarshal(body, &equipment)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [RegisterEquipment]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleEquipment) updateEquipment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var equipment entities.Equipment
	err = json.Unmarshal(body, &equipment)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.UpdateEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [UpdateEquipment]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleEquipment) deleteEquipment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	equipment := entities.Equipment{
		ID: int64(id),
	}

	err = m.useCase.DeleteEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [DeleteEquipment]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleEquipment) detailsEquipment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := entities.Equipment{
		ID: int64(id),
	}

	categoryDetails, err := m.useCase.DetailsEquipment(ctx, category)
	if err != nil {
		log.Printf("Error in [DetailsEquipment]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(categoryDetails)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
