package modules

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Module interface {
	Name() string
	Path() string
	Setup(r *mux.Router) *mux.Router
}

type ModuleHandler struct {
	Path    string
	Label   string
	Handler http.HandlerFunc
	Methods []string
}
