package controllers

import "net/http"

func RegisterControllers() {
	controller := newUserController()

	// the end slash is actually a different route in go
	http.Handle("/users", *controller)
	http.Handle("/users/", *controller)
}