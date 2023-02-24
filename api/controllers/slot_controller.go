package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zineb-ada/cyrkl/api/middlewares"
	"github.com/zineb-ada/cyrkl/api/models"
	"github.com/zineb-ada/cyrkl/api/responses"
	"github.com/zineb-ada/cyrkl/api/utils/formaterror"
)

func (server *Server) CreateSlot(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date := models.Slot{}
	err = json.Unmarshal(body, &date)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date.PrepareSlot()
	err = date.ValidateSlot()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != date.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	dateCreated, err := date.SaveSlot(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, dateCreated.ID))
	responses.JSON(w, http.StatusCreated, dateCreated)
}

func (server *Server) GetSlots(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	calendar := models.Slot{}

	calendars, err := calendar.FindSlots(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, calendars)
}

func (server *Server) GetSlotByID(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	date := models.Slot{}

	dateReceived, err := date.FindSlotByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dateReceived)
}

func (server *Server) GetUsersSlotsByUserID(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	dates := models.Slot{}
	datesReceived, err := dates.FindSlotsByUserID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

func (server *Server) UpdateSlot(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
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
	date := models.Slot{}
	err = json.Unmarshal(body, &date)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	date.PrepareSlot()
	date.ID = uint32(cid)
	err = date.ValidateSlot()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != date.UserID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	dateUpdated, err := date.UpdateASlot(server.DB, cid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, dateUpdated)
}

func (server *Server) DeleteSlot(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	date := models.Slot{}
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	userID := date.UserID
	_, err = date.DeleteASlot(server.DB, cid, userID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("user %d has deleted date %d", userID, cid))
	responses.JSON(w, http.StatusNoContent, "")
}
