package db

import (
	"database/sql"
	"fmt"
)

func getUserByQuery(query *sql.Stmt, data string) (*User, error) {
	rows, err := query.Query(data)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("user not found")
	}

	user := User{}
	err = rows.Scan(
		&user.Username,
		&user.IsEditor,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByAPIKey(apiKey string) (*User, error) {
	query, err := GetStatement("GetUserByAPIKey")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	user, err := getUserByQuery(query, apiKey)
	if err != nil {
		return nil, err
	}

	user.ApiKey = apiKey
	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	query, err := GetStatement("GetUserByUsername")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	user, err := getUserByQuery(query, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
