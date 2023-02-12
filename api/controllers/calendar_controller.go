package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zineb-ada/cyrkl/api/auth"
	"github.com/zineb-ada/cyrkl/api/models"
	"github.com/zineb-ada/cyrkl/api/responses"
	"github.com/zineb-ada/cyrkl/api/utils/formaterror"
)

func (server *Server) CreateDate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date := models.Calendar{}
	err = json.Unmarshal(body, &date)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date.PrepareCalendar()
	err = date.ValidateCalendar()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != date.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	dateCreated, err := date.SaveDate(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, dateCreated.ID))
	responses.JSON(w, http.StatusCreated, dateCreated)
}

func (server *Server) GetCalendar(w http.ResponseWriter, r *http.Request) {

	calendar := models.Calendar{}

	calendars, err := calendar.FindAllCalendar(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, calendars)
}

func (server *Server) GetDate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	date := models.Calendar{}

	dateReceived, err := date.FindDateByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dateReceived)
}

func (server *Server) GetUsersCalendarByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	dates := models.Calendar{}
	datesReceived, err := dates.FindCalendarsByUserID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

func (server *Server) UpdateDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date := models.Calendar{}
	err = json.Unmarshal(body, &date)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date.PrepareCalendar()
	date.ID = uint32(cid)
	err = date.ValidateCalendar()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != date.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	dateUpdated, err := date.UpdateADate(server.DB, cid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, dateUpdated)
}

func (server *Server) DeleteDate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	date := models.Calendar{}

	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	userID := date.UserID
	_, err = date.DeleteADate(server.DB, cid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("user %d has deleted date %d", userID, cid))
	responses.JSON(w, http.StatusNoContent, "")
}
