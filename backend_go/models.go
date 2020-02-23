package main

// Auth model
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User model
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Modified  string `json:"modified"`
	Created   string `json:"created"`
	Admin     int    `json:"admin"`
}
