package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// authentication of JWT token for each request
func jwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		// when trying to register or authenticate a token is not expected
		if r.RequestURI == "/register" || r.RequestURI == "/authenticate" {
			next.ServeHTTP(w, r)
			return
		}

		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("jwt.key")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		jwtMiddleware.HandlerWithNext(w, r, next.ServeHTTP)
	})
}

// handling authentication request, returning JWT
func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var auth Auth
	json.Unmarshal(body, &auth)

	userAuthenticated, role := authenticateUser(auth.Username, auth.Password)

	if userAuthenticated {
		generateUserCredentials(w, auth.Username, role)
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

// checking if user exists for given credentials returning 'admin' flag
func authenticateUser(username string, password string) (bool, int) {
	var user User
	db := dbConnection()
	row := db.QueryRow("SELECT admin FROM users WHERE email = ? AND password = md5(?)", username, password)
	err := row.Scan(&user.Admin)
	defer db.Close()

	if err != nil {
		return false, -1
	}

	return true, user.Admin
}

// fetching list of all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db := dbConnection()
	result, err := db.Query("SELECT id, email, firstname, lastname, modified, created, admin FROM users")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var user User
		var firstname sql.NullString
		err = result.Scan(&user.ID, &user.Email, &firstname, &user.Lastname, &user.Modified, &user.Created, &user.Admin)

		if err != nil {
			panic(err.Error())
		}

		if firstname.Valid {
			user.Firstname = firstname.String
		}

		users = append(users, user)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// creating a new user, returning JWT
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)

	if err != nil {
		panic(err.Error())
	}

	db := dbConnection()
	_, err1 := db.Exec("INSERT INTO users (email, firstname, lastname, password) values (?, ?, ?, md5(?))",
		user.Email,
		user.Firstname,
		user.Lastname,
		user.Password)

	defer db.Close()

	if err1 != nil {
		panic(err1.Error())
	}

	generateUserCredentials(w, user.Email, 0)
}

// updating existing user data
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(body, &user)

	var err error
	db := dbConnection()

	if user.Password != "" {
		_, err = db.Exec("UPDATE users SET email = ?, firstname = ?, lastname = ?, password = md5(?), modified = now() WHERE id = ?",
			user.Email,
			user.Firstname,
			user.Lastname,
			user.Password,
			user.ID)
	} else {
		// skipping password update if password is not provided
		_, err = db.Exec("UPDATE users SET email = ?, firstname = ?, lastname = ?, modified = now() WHERE id = ?",
			user.Email,
			user.Firstname,
			user.Lastname,
			user.ID)
	}

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
}

// deleting existing user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	db := dbConnection()
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
}
