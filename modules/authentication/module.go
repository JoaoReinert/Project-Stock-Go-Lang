package authentication

import (
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/modules"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type moduleAuthentication struct {
	config  entities.Config
	useCase domain.AuthenticationUseCase
	name    string
	path    string
}

func NewAuthenticationModule(config entities.Config, usecase domain.AuthenticationUseCase) modules.Module {
	return &moduleAuthentication{
		config:  config,
		useCase: usecase,
		name:    "Authentication module",
		path:    "/auth",
	}
}

func (m moduleAuthentication) Name() string {
	return m.name
}

func (m moduleAuthentication) Path() string {
	return m.path
}

func (m moduleAuthentication) Setup(r *mux.Router) *mux.Router {
	handlers := []modules.ModuleHandler{
		{
			Handler: m.registerUser,
			Path:    "/register",
			Label:   "Register a new user in database",
			Methods: []string{http.MethodPost},
		},
		{
			Handler: m.loginUser,
			Path:    "/login",
			Label:   "Login user in server",
			Methods: []string{http.MethodPost},
		},
	}

	for _, h := range handlers {
		r.HandleFunc(h.Path, h.Handler).Methods(h.Methods...)
	}

	api := r.PathPrefix("/api").Subrouter()

	api.Use(m.sessionMiddleware)

	return api
}

func (m *moduleAuthentication) sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()

		authHeader := r.Header.Get("Authorization")

		var user *entities.User
		var err error

		var token string
		if authHeader != "" {
			token = strings.ReplaceAll(authHeader, "Bearer ", "")
		}

		if token == "" {
			log.Printf("No token found in the request")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err = m.useCase.CheckDefaultSecurityToken(c, entities.UserToken{Token: token})
		if err != nil {
			log.Printf("Error in [CheckDefaultSecurityToken]")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m moduleAuthentication) registerUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
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

func (m moduleAuthentication) loginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in [ReadAll]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userCredentials entities.UserLogin

	err = json.Unmarshal(body, &userCredentials)
	if err != nil {
		log.Printf("Error in [Unmarshal]: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := m.useCase.CheckUserCredentials(ctx, userCredentials)
	if err != nil {
		log.Printf("Error in [CheckUserCredentials]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Login or password incorrect", http.StatusInternalServerError)
		return
	}

	token, err := m.useCase.GenerateTokenUser(*user)
	if err != nil {
		log.Printf("Error in [GenerateTokenUser]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsBytes, err := json.Marshal(token)
	if err != nil {
		log.Printf("Error in [json.Marshal]: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsBytes)
}
