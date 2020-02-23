package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// setting configuration from config file
func setConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("config file error: %s", err))
	}
}

// enabling CORS (for production origin and headers should be set properly)
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

// returning DB connection
func dbConnection() (db *sql.DB) {
	dbDriver := viper.GetString("database.driver")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	dbHost := viper.GetString("database.host")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbHost+"/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// generating JWT
func generateToken(username string, admin int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  username,
		"admin": admin,
		"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":   time.Now().Unix(),
	})

	return token.SignedString([]byte(viper.GetString("jwt.key")))
}

// generating user credentials (token, username and admin flag)
func generateUserCredentials(w http.ResponseWriter, username string, admin int) {
	tokenString, err := generateToken(username, admin)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"token_generation_failed"}`))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + tokenString + `", "username":"` + username + `", "admin":` + strconv.Itoa(admin) + `}`))
}
