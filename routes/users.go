package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/middleware"
	"waysbucks/repositories"
	"waysbucks/pkg/mysql"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", middleware.Auth(h.FindUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}
