package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"githubn.com/Germanicus1/lenslocked/context"
	"githubn.com/Germanicus1/lenslocked/models"
)

type Users struct {
	Templates struct {
		New            Template
		SignIn         Template
		ForgotPassword Template
		CheckYourEmail Template
		ResetPassword  Template
	}
	UserService          *models.UserService
	SessionService       *models.SessionService
	PasswordResetService *models.PasswordResetService
	MagicLinkService     *models.MagicLinkService
	EmailService         *models.EmailService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a user")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "UserService.Create: Something went wrong", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println("SessionService.Create:", err)
		// TODO: Long term, we should show a warning about not being able to sign the user in.
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
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
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	deleteCookie(w, CookieSession)
	http.Redirect(w, r, "/signin", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	user := context.User(r.Context())
	// if user == nil {
	// 	http.Redirect(w, r, "/signin", http.StatusFound)
	// 	return
	// }
	fmt.Fprintf(w, "Current user: %s\n", user.Email)
}

func (u Users) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email            string
		magicLinkChecked bool
	}
	data.Email = r.FormValue("email")
	_, data.magicLinkChecked = r.Form["magiclink"]
	u.Templates.ForgotPassword.Execute(w, r, data)
}

func (u Users) ProcessForgotPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
		Vals  url.Values
		Route string
	}
	data.Email = r.FormValue("email")
	// Checking for magic link checkbox. If the magic link
	// checkbox is not cheched, isChecked is nil.
	_, isChecked := r.Form["magiclink"]
	if isChecked {
		fmt.Println("Magic link checkbox checked")
		ml, err := u.MagicLinkService.CreateMagicLink(data.Email)
		if err != nil {
			fmt.Println("ProcessForgotPassword.CreateMagicLink:", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
			return
		}
		data.Vals = url.Values{
			"mltoken": {ml.MLToken},
		}
		data.Route = "mlsignin"
	} else {
		fmt.Println("Magic link checkbox NOT checked")

		pwReset, err := u.PasswordResetService.Create(data.Email)
		if err != nil {
			// TODO: Handle other error cases in the future (f.ex. email doesn't exist)
			fmt.Println("ProcessForgotPassword.Create:", err)
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
			return
		}
		// TODO: "" needs to be the full path with domain.
		data.Vals = url.Values{
			"token": {pwReset.Token},
		}
		data.Route = "reset-pw"
	}
	// fmt.Println("vals:", data.Vals)
	resetURL := "http://localhost:3000/" + data.Route + "?" + data.Vals.Encode()
	fmt.Println(resetURL)
	err := u.EmailService.ForgotPassword(data.Email, resetURL)
	if err != nil {
		// TODO: Handle other error cases in the future (f.ex. email doesn't exist)
		fmt.Println("ProcessForgotPassword.ForgotPassword:", err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	u.Templates.CheckYourEmail.Execute(w, r, data)
}

type UserMiddleware struct {
	SessionService *models.SessionService
}

func (umw UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// First try to read the cookie. If we run into an error reading it,
		// proceed with the request. The goal of this middleware isn't to limit
		// access. It only sets the user in the context if it can.
		token, err := readCookie(r, CookieSession)
		if err != nil {
			// fmt.Println("SetUser: Token is empty")
			// Cannot lookup the user with no cookie, so proceed without a user being
			// set, then return.
			next.ServeHTTP(w, r)
			return
		}
		// If we have a token, try to lookup the user with that token.
		user, err := umw.SessionService.User(token)
		if err != nil {
			// Invalid or expired token. In either case we can still proceed, we just
			// cannot set a user.
			next.ServeHTTP(w, r)
			return
		}
		// If we get to this point, we have a user that we can store in the context!
		// Get the context
		ctx := r.Context()
		// We need to derive a new context to store values in it. Be certain that
		// we import our own context package, and not the one from the standard
		// library.
		ctx = context.WithUser(ctx, user)
		// Next we need to get a request that uses our new context. This is done
		// in a way similar to how contexts work - we call a WithContext function
		// and it returns us a new request with the context set.
		r = r.WithContext(ctx)
		// Finally we call the handler that our middleware was applied to with the
		// updated request.
		next.ServeHTTP(w, r)
	})
}

func (u Users) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token string
	}
	data.Token = r.FormValue("token")
	u.Templates.ResetPassword.Execute(w, r, data)
}

func (u Users) ProcessMagicLink(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token string
	}
	data.Token = r.URL.Query().Get("mltoken")

	fmt.Println("ProcessMagicLink: data.Token: ", data.Token)
	user, err := u.MagicLinkService.ConsumeMagicLink(data.Token)
	if err != nil {
		fmt.Println("ProcessMagicLink.Consume:", err)
		// TODO: Distinguish between differnt types of errors
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	// Sign in the user now that the password has been reset.
	// Any errors from this point onwards should redirect the user to the sign
	// in  page.
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println("SessionService.Create:", err)
		// TODO: Distinguish between differnt types of errors
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)

}

func (u Users) ProcessResetPassword(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Token    string
		Password string
	}
	data.Token = r.FormValue("token")
	data.Password = r.FormValue("password")
	user, err := u.PasswordResetService.Consume(data.Token)
	if err != nil {
		fmt.Println("ProcessResetPassword.Consume:", err)
		// TODO: Distinguish between differnt types of errors
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = u.UserService.UpdatePassword(user.ID, data.Password)
	if err != nil {
		fmt.Println("ProcessResetPassword.UpdatePassword:", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// Sign in the user now that the password has been reset.
	// Any errors from this point onwards should redirect the user to the sign
	// in  page.
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println("SessionService.Create:", err)
		// TODO: Distinguish between differnt types of errors
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (umw UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		if user == nil {
			fmt.Println("No user logged in")
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
