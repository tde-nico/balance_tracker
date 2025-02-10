package db

import "fmt"

func GetTotalBalance() (float32, float32, error) {
	query, err := GetStatement("GetTotalBalance")
	if err != nil {
		return 0, 0, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return 0, 0, fmt.Errorf("error querying balance: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, 0, fmt.Errorf("error scanning balance: %v", err)
	}

	var balance, fakeBalance float32
	err = rows.Scan(&balance, &fakeBalance)
	if err != nil {
		return 0, 0, fmt.Errorf("error scanning balance: %v", err)
	}
	return balance, fakeBalance, nil
}
