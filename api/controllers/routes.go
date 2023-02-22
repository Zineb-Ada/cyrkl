package controllers

import "github.com/zineb-ada/cyrkl/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/updateuser/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("POST")
	s.Router.HandleFunc("/deleteuser/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// calendar routes
	s.Router.HandleFunc("/slot", middlewares.SetMiddlewareJSON(s.CreateSlot)).Methods("POST")
	s.Router.HandleFunc("/slots", middlewares.SetMiddlewareJSON(s.GetSlots)).Methods("GET")
	s.Router.HandleFunc("/slot/{id}", middlewares.SetMiddlewareJSON(s.GetSlotByID)).Methods("GET")
	s.Router.HandleFunc("/slots/user/{user_id}", middlewares.SetMiddlewareJSON(s.GetUsersSlotsByUserID)).Methods("GET")
	s.Router.HandleFunc("/updateslot/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSlot))).Methods("PUT")
	s.Router.HandleFunc("/deleteslot/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSlot)).Methods("DELETE")

	// Invitations routes
	s.Router.HandleFunc("/invitation", middlewares.SetMiddlewareJSON(s.CreateInvitation)).Methods("POST")
	s.Router.HandleFunc("/invitations", middlewares.SetMiddlewareJSON(s.GetInvitions)).Methods("GET")
	s.Router.HandleFunc("/invitation/{id}", middlewares.SetMiddlewareJSON(s.GetInvitationByID)).Methods("GET")
	s.Router.HandleFunc("/invitationsreceived/user/{invited_id}", middlewares.SetMiddlewareJSON(s.GetInvitationsReceivedByInvitedID)).Methods("GET")
	s.Router.HandleFunc("/invitationssended/user/{inviter_id}", middlewares.SetMiddlewareJSON(s.GetInvitationsSendedByInviterID)).Methods("GET")
	s.Router.HandleFunc("/updateinvitation/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateInvitation))).Methods("PUT")
	s.Router.HandleFunc("/deleteinvitation/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDate)).Methods("DELETE")

	// routes de la liste des Users triés par ordre de pertinence par l'algorithme
	s.Router.HandleFunc("/usersalgo/{id}", middlewares.SetMiddlewareJSON(s.GetInvitions)).Methods("GET")

}
