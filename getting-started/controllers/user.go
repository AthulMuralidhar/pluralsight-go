package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"pluralsight-go/getting-started/models"
	"regexp"
	"strconv"
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
	if req.URL.Path == "/users" {
		switch req.Method {
		case http.MethodGet:
			controller.getAll(writer, req)
		case http.MethodPost:
			controller.post(writer, req)
		default:
			writer.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// single user struct handelling branch
		matches := controller.userIdPattern.FindStringSubmatch(req.URL.Path)

		if len(matches) == 0 {
			writer.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
		}
		switch req.Method {
		case http.MethodGet:
			controller.get(id, writer)
		case http.MethodPost:
			controller.put(id, writer, req)
		case http.MethodDelete:
			controller.delete(id, writer)
		default:
			writer.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func (controller *userController) getAll(writer http.ResponseWriter, req *http.Request) {
	encodeResponseJSON(models.GetAllUsers(), writer)
}

func (controller *userController) get(id int, writer http.ResponseWriter) {
	usr, err := models.GetUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("cant find user"))
		return
	}
	encodeResponseJSON(usr, writer)

}

func (controller *userController) post(writer http.ResponseWriter, req *http.Request) {
	// runtime.Breakpoint()

	usr, err := controller.parseRequest(req)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not parse user"))
		return
	}
	usr, err = models.AddUser(usr)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("erroring while adding user"))
		// fmt.Errorf("%v",err)
		return
	}
	encodeResponseJSON(usr, writer)

}

func (controlelr *userController) put(id int, writer http.ResponseWriter, req *http.Request) {
	usr, err := controlelr.parseRequest(req)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("cannot parse user"))
	}
	usr, err = models.UpdateUser(usr)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("cannot retrieve user"))
	}
	encodeResponseJSON(usr, writer)
}

func (controller *userController) delete(id int, writer http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("cannot delete user"))
	}
	writer.WriteHeader(http.StatusOK)
}

func (controller *userController) parseRequest(req *http.Request) (models.User, error) {
	decoder := json.NewDecoder(req.Body)
	var usr models.User
	err := decoder.Decode(&usr)

	if err != nil {
		return models.User{}, err
	}

	return usr, nil
}

func encodeResponseJSON(data interface{}, writer io.Writer) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(data)
}
