package api

import (
	"net/http"

	"chit-chat/internal/di/util"
	"chit-chat/internal/services/auth"
)

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := auth.CheckSession(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		util.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}
