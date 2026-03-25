package category

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

type moduleCategory struct {
	useCase domain.CategoryUseCase
	path    string
	name    string
}

func NewCategoryModule(
	useCase domain.CategoryUseCase,
) modules.Module {
	return moduleCategory{
		useCase: useCase,
		name:    "Category module",
		path:    "/category",
	}
}

func (m moduleCategory) Name() string {
	return m.name
}

func (m moduleCategory) Path() string {
	return m.path
}

func (m moduleCategory) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerCategory,
			Path:    "/register",
			Label:   "Register category in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.updateCategory,
			Path:    "/update",
			Label:   "Update category in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.deleteCategory,
			Path:    "/delete/{id}",
			Label:   "Delete category in database",
			Methods: []string{http.MethodDelete},
		},
		{
			Handler: m.detailsCategory,
			Path:    "/details/{id}",
			Label:   "Get details category in database",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleCategory) registerCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category entities.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleCategory) updateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var category entities.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.UpdateCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [UpdateCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleCategory) deleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := entities.Category{
		ID: int64(id),
	}

	err = m.useCase.DeleteCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [DeleteCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleCategory) detailsCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := entities.Category{
		ID: int64(id),
	}

	categoryDetails, err := m.useCase.DetailsCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [DeleteCategory]")
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
