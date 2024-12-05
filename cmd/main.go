package main

import (
    "log"
    "net/http"
    "user-management/config"
    "user-management/routes"
    "user-management/utils"

    "github.com/gorilla/mux"
)

func main() {
    cfg := config.GetConfig()
    r := mux.NewRouter()

    // Log middleware untuk setiap request
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            utils.LogRequest(r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
        })
    })

    routes.RegisterRoutes(r)

    log.Printf("Server running on port %s", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
