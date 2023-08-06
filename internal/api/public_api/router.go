package public_api

import (
	"net/http"

	"chit-chat/internal/di/config"
	thread_api "chit-chat/internal/entities/threads/api"
	"chit-chat/internal/services/auth"
)

func NewRouter(cfg *config.Configuration) *http.ServeMux {
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(cfg.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/logout", auth.Logout)
	mux.HandleFunc("/signup", auth.Signup)
	mux.HandleFunc("/signup_account", auth.SignupAccount)
	mux.HandleFunc("/authenticate", auth.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", thread_api.NewThread)
	mux.HandleFunc("/thread/create", thread_api.CreateThread)
	mux.HandleFunc("/thread/post", thread_api.PostThread)
	mux.HandleFunc("/thread/read", thread_api.ReadThread)

	return mux
}
