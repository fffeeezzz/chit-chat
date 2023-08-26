package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"chit-chat/internal/di/util"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	t := util.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	util.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		util.Danger(&log.Logger{}, err, "Cannot parse form")
	}
	user := User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		util.Danger(&log.Logger{}, err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := UserByEmail(request.PostFormValue("email"))
	if err != nil {
		util.Danger(&log.Logger{}, err, "Cannot find user")
	}
	if user.Password == util.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			util.Danger(&log.Logger{}, err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		util.Warning(&log.Logger{}, err, "Failed to get cookie")
		session := Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}

var errInvalidSession = errors.New("invalid session")

// Checks if the user is logged in and has a session, if not err is not nil
func CheckSession(_ http.ResponseWriter, request *http.Request) (*Session, error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		return nil, fmt.Errorf("get cookie: %w", err)
	}

	session := Session{Uuid: cookie.Value}
	if ok, _ := session.Check(); !ok {
		return nil, errInvalidSession
	}

	return &session, nil
}
