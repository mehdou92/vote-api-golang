package controllers

import "github.com/mehdou92/vote-api/api/middlewares"

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

	//Votes routes
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.CreateVote)).Methods("POST")
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.GetVotes)).Methods("GET")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareJSON(s.GetVote)).Methods("GET")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateVote))).Methods("PUT")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteVote)).Methods("DELETE")
}
