package api

import (
	"net/http"

	"chit-chat/internal/di/util"
	"chit-chat/internal/entities/threads/domain"
	"chit-chat/internal/services/auth"
)

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := domain.ThreadByUUID(uuid)
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot read thread")
	} else {
		_, err := auth.CheckSession(writer, request)
		if err != nil {
			util.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			util.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}
