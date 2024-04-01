package controllers

import (
	"lenslocked/views"
	"net/http"
)

func StaticHandler(template views.Template, data any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, data)
	}
}
