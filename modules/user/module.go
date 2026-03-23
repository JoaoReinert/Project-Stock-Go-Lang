package user

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type moduleUser struct {
	useCase domain.UserUseCase
	name    string
	path    string
}

func NewUserModule(useCase domain.UserUseCase) modules.Module {
	return &moduleUser{
		useCase: useCase,
		name:    "User module",
		path:    "/user",
	}
}

func (m moduleUser) Name() string {
	return m.name
}

func (m moduleUser) Path() string {
	return m.path
}

func (m moduleUser) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerUser,
			Path:    "/register",
			Label:   "Register a new user in database",
			Methods: []string{http.MethodPost},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	return r
}

func (m moduleUser) registerUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user entities.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.useCase.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Error in [RegisterUser]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
