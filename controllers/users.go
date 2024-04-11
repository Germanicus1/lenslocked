package controllers

import (
	"fmt"
	"net/http"

	"githubn.com/Germanicus1/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Somethiong went wrong.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "cookie set: %+v\n", cookie)
	fmt.Fprintf(w, "user authenticated: %+v\n", user)
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Somethiong went wrong.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created: %+v", user) // +v gives fieldnames as well
}
