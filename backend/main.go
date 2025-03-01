package main

import (
    "log"
    "net/http"
    "portfolio/config"
    "portfolio/routes"

    "github.com/gorilla/mux"
)

func main() {
    config.ConnectDB()

    router := mux.NewRouter()
    routes.RegisterProjectRoutes(router)

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
