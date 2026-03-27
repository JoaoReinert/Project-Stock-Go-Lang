package equipment

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
			Label:   "Register new equipment in database",
			Methods: []string{http.MethodPost},
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
		log.Printf("Error in [ReadAll] %v", err)
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
