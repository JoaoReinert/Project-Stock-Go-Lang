package statistics

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type moduleStatistics struct {
	useCase domain.StatisticsUseCase
	path    string
	name    string
}

func NewStatisticsModule(
	useCase domain.StatisticsUseCase,
) modules.Module {
	return moduleStatistics{
		useCase: useCase,
		name:    "Statistics module",
		path:    "/statistics",
	}
}

func (m moduleStatistics) Name() string {
	return m.name
}

func (m moduleStatistics) Path() string {
	return m.path
}

func (m moduleStatistics) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.getTotalNumberEntriesPerDate,
			Path:    "/number/entries",
			Label:   "Return total number of entries in a per period",
			Methods: []string{http.MethodGet},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleStatistics) getTotalNumberEntriesPerDate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	startDate := r.URL.Query().Get("starDate")
	endDate := r.URL.Query().Get("endDate")

	startDateConvert, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Printf("error in time.Parse (startDate): %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endDateConvert, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Printf("error in time.Parse (endDate): %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listEquipments, err := m.useCase.GetTotalNumberPerDate(ctx, startDateConvert, endDateConvert, true)
	if err != nil {
		log.Printf("Error in [getTotalNumberEntriesPerDate]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(listEquipments)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response)
}
