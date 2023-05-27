package models

import (
	"github.com/gustablo/secret-auth/config"
	"github.com/gustablo/secret-auth/internal/commons"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) Insert(ownerId string) error {
	commons.Generate(&u.ID)
	commons.Encrypt(&u.Password)

	_, err := config.Conn.Exec("INSERT INTO users (id, username, email, password_hash, owner_id) VALUES (?, ?, ?, ?, ?)",
		u.ID, u.Username, u.Email, u.Password, ownerId)

	if err != nil {
		return err
	}

	return nil
}

func (u User) SelectByEmailOrUserName(ownerId string) (User, error) {
	var user User = User{}

	rows, err := config.Conn.Query("SELECT * FROM users WHERE (email = ? OR username = ?) AND owner_id = ?", u.Email, u.Username, ownerId)
	if err != nil {
		return user, err
	}

	defer rows.Close()

	rows.Next()

	if err := rows.Err(); err != nil {
		return user, err
	}

	err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &ownerId)

	if err != nil {
		return user, err
	}

	return user, nil
}
