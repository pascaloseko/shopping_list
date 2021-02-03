package model

import (
	"github.com/pascaloseko/shopping_list/server/database"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID           uint64 `json:"id"`
	UserName     string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisterDate string `json:"register_date"`
}

// Item struct
type Item struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

// Items struct
type Items struct {
	Items []Item `json:"items"`
}

func (user *User) UserRegister() (response map[string]interface{}, err error) {
	// insert user to db
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) returning id, name, email, password"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	if err := stmt.QueryRow(user.UserName, user.Email, Encrypt(user.Password)).Scan(&user.ID, &user.UserName, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	response = map[string]interface{}{"id": user.ID, "name": user.UserName, "email": user.Email, "register_date": user.RegisterDate}
	return
}

// Encrypt encypts a string with sha1 algorithm
func Encrypt(plaintext string) (cryptoText string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	cryptoText = string(hashPassword)
	return
}
