package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func newUserController() *userController {
	//  the addressof operator "&" can be used directly for structs
	// this is not allowed for primitive data types as there is no
	// memory allocated without a variable declaration present


	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}

}

// checking in the net/http package for 
// https://pkg.go.dev/net/http#Handler
// 
// the handler interface has the same name as our func below so that is enough
// for go to recognize that there is an interface at work here
func (controller userController) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("user controller is working"))
}