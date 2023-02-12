package controllers

import "github.com/zineb-ada/cyrkl/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// calendar routes
	s.Router.HandleFunc("/calendar", middlewares.SetMiddlewareJSON(s.CreateDate)).Methods("POST")
	s.Router.HandleFunc("/calendar", middlewares.SetMiddlewareJSON(s.GetCalendar)).Methods("GET")
	s.Router.HandleFunc("/calendar/{id}", middlewares.SetMiddlewareJSON(s.GetDate)).Methods("GET")
	s.Router.HandleFunc("/calendar/user/{user_id}", middlewares.SetMiddlewareJSON(s.GetUsersCalendarByUserID)).Methods("GET")
	s.Router.HandleFunc("/calendar/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateDate))).Methods("PUT")
	s.Router.HandleFunc("/calendar/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDate)).Methods("DELETE")

}
