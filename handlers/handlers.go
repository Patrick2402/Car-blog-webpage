package handlers

import "net/http"

func HomeHandler(write http.ResponseWriter, read *http.Request) {
	http.ServeFile(write, read, "./html/index.html")
}

func LoginHandler(write http.ResponseWriter, read *http.Request) {
	http.ServeFile(write, read, "./html/login.html")
}

func BlogHandler(write http.ResponseWriter, read *http.Request) {
	http.ServeFile(write, read, "./html/blog.html")
}
