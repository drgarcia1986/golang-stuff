package main

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateList map[string]*template.Template

type TemplateHandler struct {
	Templates  TemplateList
	HandleFunc func(w http.ResponseWriter, r *http.Request, templates TemplateList)
}

func (th TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.HandleFunc(w, r, th.Templates)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, templates TemplateList) {
	queryString := r.URL.Query()
	name, ok := queryString["name"]
	if !ok {
		name = []string{"Stranger"}
	}
	err := templates["index"].ExecuteTemplate(
		w, "base", map[string]string{"Name": name[0]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	templates := TemplateList{
		"index": template.Must(template.ParseFiles(
			"index.html", "base.html")),
	}
	http.Handle("/", TemplateHandler{Templates: templates, HandleFunc: IndexHandler})

	log.Println("Start Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
