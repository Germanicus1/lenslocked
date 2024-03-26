package controllers

import (
	"html/template"
	"net/http"

	"githubn.com/Germanicus1/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct{
		Question string
		Answer template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer: "Yes, we offer a 30 day trial",
		},
		{
			Question: "What are your support hours?",
			Answer: "We have support stuff answering email 24/7",
		},
		{
			Question: "How do I contact support?",
			Answer: `Email us at <a href="mailto:peter.kerschbaumer.es">peter@kerschbaumer.es</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request){
		tpl.Execute(w, questions)
	}
}
