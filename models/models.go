package models

import (
	"time"

	"github.com/AdluAghnia/xhater/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id,omitempty"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserSession struct {
	SID    string `json:"sid"`
	IP     string `json:"ip"`
	Expiry string `json:"expiry"`
	UA     string `json:"ua"`
}

type Account struct {
	Email    string        `json:"email"`
	Username string        `json:"username"`
	Session  string        `json:"userSession"`
	Sessions []UserSession `json:"session"`
}

func (u *User) CreateUser() (User, error) {
	query := "INSERT INTO user (id, email, username, password, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	// hash the password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		return User{}, err
	}

	var newUser User
	err = stmt.QueryRow(
		uuid.NewString(),
		u.Email,
		u.Username,
		string(hashedpassword),
		time.Now().UTC(),
	).Scan(
		&newUser.ID,
		&newUser.Email,
		&newUser.Username,
		&newUser.Password,
		&newUser.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return newUser, nil

}
