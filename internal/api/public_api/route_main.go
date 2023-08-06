package public_api

import (
	"net/http"

	"chit-chat/data"
	"chit-chat/internal/di/util"
	"chit-chat/internal/services/auth"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := auth.CheckSession(writer, request)
	if err != nil {
		util.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		util.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		util.ErrorMessage(writer, request, "Cannot get threads")
	} else {
		_, err := auth.CheckSession(writer, request)
		if err != nil {
			util.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			util.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
