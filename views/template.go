package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		fmt.Println("Error while parsing template:", t)
		panic("Closing program")
	}
	return t
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	t, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing FS template: %w", err)
	}
	return Template{
		htmlTemplate: t,
	}, nil
}

func Parse(filepath string) (Template, error) {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTemplate: t,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("Error while rendering template: %v", err)
		http.Error(w, "Error while generating the page", http.StatusInternalServerError)
		return
	}
}
