package routes

import (
    "user-management/controllers"

    "github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
    r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
