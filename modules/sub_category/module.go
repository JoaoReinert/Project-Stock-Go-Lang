package sub_category

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

type moduleSubCategory struct {
	useCase domain.SubCategoryUseCase
	path    string
	name    string
}

func NewSubCategoryModule(useCase domain.SubCategoryUseCase) modules.Module {
	return moduleSubCategory{
		useCase: useCase,
		name:    "SubCategory module",
		path:    "/subCategory",
	}
}

func (m moduleSubCategory) Name() string {
	return m.name
}

func (m moduleSubCategory) Path() string {
	return m.path
}

func (m moduleSubCategory) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerSubCategory,
			Path:    "/register",
			Label:   "Register subCategory in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.updateSubCategory,
			Path:    "/update",
			Label:   "Update subCategory in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.deleteSubCategory,
			Path:    "/delete/{id}",
			Label:   "Delete subCategory in database",
			Methods: []string{http.MethodDelete},
		},
		{
			Handler: m.detailsSubCategory,
			Path:    "/details/{id}",
			Label:   "Get details subCategory in database",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleSubCategory) registerSubCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll] %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var subCategory entities.SubCategory
	err = json.Unmarshal(body, &subCategory)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleSubCategory) updateSubCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll] %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var subCategory entities.SubCategory
	err = json.Unmarshal(body, &subCategory)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.UpdateSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [UpdateSubCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleSubCategory) deleteSubCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subCategory := entities.SubCategory{
		ID: int64(id),
	}

	err = m.useCase.DeleteSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [DeleteSubCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m moduleSubCategory) detailsSubCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("Error in [Atoi]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subCategory := entities.SubCategory{
		ID: int64(id),
	}

	subCategoryDetails, err := m.useCase.DetailsSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [DetailsSubCategory]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(subCategoryDetails)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
