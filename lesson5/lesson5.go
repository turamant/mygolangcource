// 5. Templates
// Добавим поддержку шаблонов для генерации HTML.


package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{
		Message: "Hello, Templates!",
	}
	renderTemplate(w, "index.html", data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8000", nil)
}
