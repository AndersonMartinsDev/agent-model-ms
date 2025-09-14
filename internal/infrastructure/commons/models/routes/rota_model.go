package model

import "net/http"

type RouteModel struct {
	URI          string
	Metodo       string
	Func         func(w http.ResponseWriter, r *http.Request)
	Autenticacao bool
}
