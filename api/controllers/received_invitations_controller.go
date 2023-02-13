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

func (server *Server) CreateReceivedInvitation(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation := models.InvitationsReceived{}
	err = json.Unmarshal(body, &invitation)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	invitation.PrepareReceivedInvit()
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != invitation.UserSenderID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	invitationCreated, err := invitation.SaveReceivedInvitation(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, invitationCreated.ID))
	responses.JSON(w, http.StatusCreated, invitationCreated)
}

func (server *Server) GetRdInvitions(w http.ResponseWriter, r *http.Request) {
	invitationsReceived := models.InvitationsReceived{}

	allinvitationsReceived, err := invitationsReceived.FindAllInvitationsReceived(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, allinvitationsReceived)
}

func (server *Server) GetRdInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	irid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	rdinvitation := models.InvitationsReceived{}
	dateReceived, err := rdinvitation.FindRdInvitationByID(server.DB, irid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dateReceived)
}

func (server *Server) GetUsersRdInvitationByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["user_receiver_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		responses.ERROR(w, http.StatusBadRequest, errors.New("missing user_id in the request"))
		return
	}
	invitationsreceived := models.InvitationsReceived{}
	datesReceived, err := invitationsreceived.FindRdInvitByUserID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datesReceived)
}


// fonctions à rajouter : getInvitationbyrefusé et getInvitbyaccepté et peut être getInvitationbyrefuséet accepté
// getInvitationbyacceptébyuserid; getInvitationbyrefusébyuserid
// invitations recues et envoyées (refusé et accepté) GetInvitationsReceivedAndSendedAccepté GetInvitationsRAndSrefusé par user

