package models

import "github.com/gustablo/secret-auth/config"

type Owner struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	AppName string `json:"app_name"`
}

func (o Owner) SelectByKey() (string, error) {
	rows, err := config.Conn.Query("SELECT id FROM owners WHERE access_key = ?", o.Key)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var id string

	rows.Next()

	if err := rows.Err(); err != nil {
		return "", err
	}

	rows.Scan(&id)

	return id, nil
}
