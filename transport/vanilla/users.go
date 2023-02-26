package vanilla

import (
	"net/http"
	"regexp"

	"vse-course/service"
	svcmodel "vse-course/service/model"
	"vse-course/transport/model"
	"vse-course/transport/util"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	//validate request
}

var usersRegExp = regexp.MustCompile(`/users/(?P<email>[\w-\.]+@([\w-]+\.)+[\w-]{2,4})`)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	match := usersRegExp.FindStringSubmatch(r.URL.Path)

	var email string
	if len(match) > 0 {
		for i, name := range usersRegExp.SubexpNames() {
			if name == "email" {
				email = match[i]
			}
		}
	} else {
		util.WriteErrResponse(w, http.StatusNotFound, nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetUser(w, r, email)
		return
	case http.MethodPatch:
		UpdateUser(w, r, email)
		return
	case http.MethodDelete:
		DeleteUser(w, r, email)
	default:
		util.WriteErrResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, email string) {
	user, err := service.GetUser(r.Context(), email)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, email string) {
	var user model.User
	err := util.UnmarshalRequest(r, &user)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := service.UpdateUser(r.Context(), email, svcmodel.User{})
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, email string) {
	err := service.DeleteUser(r.Context(), email)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusNoContent, nil)
}
