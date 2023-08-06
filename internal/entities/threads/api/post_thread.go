package api

import (
	"log"
	"net/http"

	"chit-chat/internal/di/util"
	"chit-chat/internal/services/auth"
)

// POST /thread/create
// Create new thread
func CreateThread(writer http.ResponseWriter, request *http.Request) {
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
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			util.Danger(&log.Logger{}, err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}
