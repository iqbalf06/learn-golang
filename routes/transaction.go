package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/middleware"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransaction)).Methods("GET")
	// r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
}