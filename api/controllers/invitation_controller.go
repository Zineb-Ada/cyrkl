package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zineb-ada/cyrkl/api/auth"
	"github.com/zineb-ada/cyrkl/api/middlewares"
	"github.com/zineb-ada/cyrkl/api/models"
	"github.com/zineb-ada/cyrkl/api/responses"
	"github.com/zineb-ada/cyrkl/api/utils/formaterror"
)

func (server *Server) CreateInvitation(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["uid"], 10, 64)
	fmt.Printf("uid 24 %v", uid)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation := models.Invitation{}
	err = json.Unmarshal(body, &invitation)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation.PrepareInvitation("create")
	invitation.InvitedID = uint32(uid)
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != invitation.InviterID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	if len(invitation.Statut) > 0 {
		invitation.Statut = strings.ToLower(invitation.Statut)
	}
	invitationCreated, err := invitation.SaveInvitation(server.DB, uid)
	if err != nil {
		// formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, invitationCreated.ID))
	responses.JSON(w, http.StatusCreated, invitationCreated)
}

func (server *Server) GetInvitions(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	invitation := models.Invitation{}

	invitations, err := invitation.FindInvitations(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, invitations)
}

func (server *Server) GetInvitationByID(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	inid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	invitation := models.Invitation{}
	invitationReceived, err := invitation.FindInvitatByID(server.DB, inid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, invitationReceived)
}

func (server *Server) GetInvitationsReceivedByInvitedID(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["invited_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	invitationsreceived := models.Invitation{}
	datesReceived, err := invitationsreceived.FindInvitsReceivedByInvitedID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

func (server *Server) GetInvitationsReceivedByInvitedIDWithStatus(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	fmt.Println(vars)
	userID, err := strconv.ParseUint(vars["invited_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	status := vars["status"]
	if len(status) > 0 {
		status = strings.ToLower(status)
	}
	if status != "in progress" && status != "accepted" && status != "refused" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid status value"))
		return
	}
	invitationsreceived := models.Invitation{}
	datesReceived, err := invitationsreceived.FindInvitsReceivedByInvitedIDWithStatus(server.DB, userID, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

func (server *Server) GetInvitationsSendedByInviterID(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["inviter_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	invitationsreceived := models.Invitation{}
	datesReceived, err := invitationsreceived.FindInvitSendedByInviterID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

func (server *Server) GetInvitationsSendedByInviterIDWithStatus(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)
	vars := mux.Vars(r)
	fmt.Println(vars)
	userID, err := strconv.ParseUint(vars["inviter_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	status := vars["status"]
	if len(status) > 0 {
		status = strings.ToLower(status)
	}
	if status != "in progress" && status != "accepted" && status != "refused" {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid status value"))
		return
	}
	invitationsreceived := models.Invitation{}
	datesReceived, err := invitationsreceived.FindInvitsSendedByInviterIDWithStatus(server.DB, userID, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

// Update Invitation
func (server *Server) CreateDate(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)

	vars := mux.Vars(r)
	inid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation := models.Invitation{}
	err = json.Unmarshal(body, &invitation)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation.PrepareInvitation("update")
	invitation.ID = uint32(inid)
	if len(invitation.Statut) > 0 {
		invitation.Statut = strings.ToLower(invitation.Statut)
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != invitation.InvitedID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	dateCreated, err := invitation.UpdateInvit(server.DB, inid)
	fmt.Printf("datecreated : 151 controlers %v", dateCreated)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, dateCreated)
}

func (server *Server) DeleteDate(w http.ResponseWriter, r *http.Request) {
	middlewares.EnableCors(&w)

	vars := mux.Vars(r)

	date := models.Invitation{}

	inid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	userName := date.Slotd.Userc.Name
	_, err = date.DeleteInvit(server.DB, inid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("user %s has deleted date %d", userName, inid))
	responses.JSON(w, http.StatusNoContent, "")
}
