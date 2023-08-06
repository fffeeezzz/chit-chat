package api

import (
	"fmt"
	"log"
	"net/http"

	"chit-chat/internal/di/util"
	"chit-chat/internal/entities/threads/domain"
	"chit-chat/internal/services/auth"
)

// POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := auth.CheckSession(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			util.Danger(&log.Logger{}, err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			util.Danger(&log.Logger{}, err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := domain.ThreadByUUID(uuid)
		if err != nil {
			util.ErrorMessage(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			util.Danger(&log.Logger{}, err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
