package util

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// ParseTemplateFiles parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) *template.Template {
	var (
		t     *template.Template
		files []string
	)

	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	t = template.Must(t.ParseFiles(files...))

	return t
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		log.Printf("execute template: %v", err)
	}
}

// for logging

func Info(logger *log.Logger, args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Danger(logger *log.Logger, args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func Warning(logger *log.Logger, args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
