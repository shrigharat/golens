package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("email: %s, password: %s", r.FormValue("email"), r.FormValue("password"))
	fmt.Fprint(w, "Temporary Response")
}
