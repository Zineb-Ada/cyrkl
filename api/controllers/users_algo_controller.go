package controllers

import (
	"net/http"

	"github.com/zineb-ada/cyrkl/api/models"
	"github.com/zineb-ada/cyrkl/api/responses"
)

func (server *Server) GetUsersAlgoByUserID(w http.ResponseWriter, r *http.Request) {
	ua := models.UsersAlgo{}

	usersalgo, err := ua.GetUsersAlgoByUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, usersalgo)
}
