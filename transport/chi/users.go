package chi

import (
	"net/http"

	"vse-course/service"
	svcmodel "vse-course/service/model"
	"vse-course/transport/model"
	"vse-course/transport/util"

	"github.com/go-chi/chi"
)

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")

	return email
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := util.UnmarshalRequest(r, &user)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
	}

	err = service.CreateUser(r.Context(), model.ToSvcUser(user))
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}

	util.WriteResponse(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := service.GetUser(r.Context(), getEmailFromURL(r))
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, model.ToNetUser(user))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := service.ListUsers(r.Context())

	util.WriteResponse(w, http.StatusOK, model.ToNetUsers(users))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := util.UnmarshalRequest(r, &user)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := service.UpdateUser(r.Context(), getEmailFromURL(r), svcmodel.User{})
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := service.DeleteUser(r.Context(), getEmailFromURL(r))
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusNoContent, nil)
}
