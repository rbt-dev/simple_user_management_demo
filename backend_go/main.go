package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	setConfig()
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/authenticate", authenticate).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/register", createUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users", getUsers).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/user", updateUser).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/user", deleteUser).Queries("id", "{id}").Methods(http.MethodDelete, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(jwtAuthentication)

	err := http.ListenAndServe(":"+viper.GetString("server.port"), router)

	if err != nil {
		fmt.Print(err)
	}
}
