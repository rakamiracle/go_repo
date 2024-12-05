package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "user-management/models"
    "user-management/storage"

    "github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users := storage.GetUsers()
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    user, err := storage.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    storage.AddUser(user)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var updatedUser models.User
    json.NewDecoder(r.Body).Decode(&updatedUser)
    updatedUser.ID = id

    err := storage.UpdateUser(id, updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    err := storage.DeleteUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
