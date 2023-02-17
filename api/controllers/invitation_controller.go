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

func (server *Server) CreateInvitation(w http.ResponseWriter, r *http.Request) {
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
	invitation.PrepareInvitation()
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != invitation.InviterID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	invitationCreated, err := invitation.SaveInvitation(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, invitationCreated.ID))
	responses.JSON(w, http.StatusCreated, invitationCreated)
}

func (server *Server) GetInvitions(w http.ResponseWriter, r *http.Request) {
	invitationsReceived := models.Invitation{}

	invitations, err := invitationsReceived.FindInvitations(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, invitations)
}

func (server *Server) GetInvitationByID(w http.ResponseWriter, r *http.Request) {
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

func (server *Server) GetInvitationByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["user_receiver_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	invitationsreceived := models.Invitation{}
	datesReceived, err := invitationsreceived.FindInvitByUserID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}

// Cette fonction est plus createDate 
func (server *Server) UpdateInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inid, err := strconv.ParseUint(vars["id"], 10, 64)
	fmt.Printf("irid 97 : %d", inid)

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
	invitation.PrepareInvitation()
	invitation.ID = uint32(inid)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if uid != invitation.InviterID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	dateCreated, err := invitation.UpdateInvit(server.DB, inid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, dateCreated)
}

func (server *Server) DeleteDate(w http.ResponseWriter, r *http.Request) {

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

// fonctions à rajouter : getInvitationbyrefusé et getInvitbyaccepté et peut être getInvitationbyrefuséet accepté
// getInvitationbyacceptébyuserid; getInvitationbyrefusébyuserid
// invitations recues et envoyées (refusé et accepté) GetInvitationsReceivedAndSendedAccepté GetInvitationsRAndSrefusé par user
