package controllers

import (
	"net/http"

	"githubn.com/Germanicus1/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tpl.Execute(w, nil)
	}
}
