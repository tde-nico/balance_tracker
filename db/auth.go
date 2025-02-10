package db

import (
	"balance/crypto_utils"
	"fmt"
)

const APIKEY_LENGTH = 32
const SALT_LENGTH = 8
const INVALID_PREFIX = "INVALID_"

func UserExists(username string) (bool, error) {
	query, err := GetStatement("UserExists")
	if err != nil {
		return false, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query(username)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

func createUser(username, salt, secret, apiKey string, isEditor bool) error {
	insert, err := GetStatement("CreateUser")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	_, err = insert.Exec(username, salt, secret, apiKey, isEditor)
	return err
}

func RegisterUser(username, password string, isEditor bool) error {
	_, apiKeyHex, err := crypto_utils.GetRand(APIKEY_LENGTH)
	if err != nil {
		return err
	}

	salt, saltHex, err := crypto_utils.GetRand(SALT_LENGTH)
	if err != nil {
		return err
	}

	secretHex := crypto_utils.HashPassword(password, salt)

	err = createUser(username, string(saltHex), string(secretHex), string(apiKeyHex), isEditor)
	return err
}

func LoginUser(username, password string) (string, error) {
	query, err := GetStatement("LoginUser")
	if err != nil {
		return "", fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query(username)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", fmt.Errorf("user \"%s\" not found", username)
	}

	var apikey, saltHex, secretHex string
	err = rows.Scan(&apikey, &saltHex, &secretHex)
	if err != nil {
		return "", err
	}

	salt, err := crypto_utils.HexToBytes(saltHex)
	if err != nil {
		return "", err
	}

	hash := crypto_utils.HashPassword(password, salt)
	if secretHex != hash {
		return "", fmt.Errorf("invalid password")
	}

	return apikey, nil
}
