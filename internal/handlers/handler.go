package handlers

import "github.com/gorilla/mux"

type Handler interface {
	Register(router *mux.Router)
}
