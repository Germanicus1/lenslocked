package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"githubn.com/Germanicus1/lenslocked/controllers"
	"githubn.com/Germanicus1/lenslocked/templates"
	"githubn.com/Germanicus1/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.Get("/about", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "about.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	fmt.Println("Starting the server at :3000...")

	http.ListenAndServe(":3000", r)
}