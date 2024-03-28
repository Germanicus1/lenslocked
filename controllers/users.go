package controllers

import (
	"net/http"

	"githubn.com/Germanicus1/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request){
	u.Templates.New.Execute(w, nil)
}