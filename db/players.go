package db

import (
	"balance/crypto_utils"
	"fmt"
)

func GetPlayers() ([]Player, error) {
	query, err := GetStatement("GetPlayers")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying players: %v", err)
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var p Player
		err = rows.Scan(&p.Username, &p.Balance, &p.FakeBalance, &p.ItemCount)
		p.Balance = crypto_utils.RoundToTwoDecimals(p.Balance)
		p.FakeBalance = crypto_utils.RoundToTwoDecimals(p.FakeBalance)
		if err != nil {
			return nil, fmt.Errorf("error scanning players: %v", err)
		}
		players = append(players, p)
	}
	return players, nil
}

func GetPlayer(username string) (*Player, error) {
	query, err := GetStatement("GetPlayer")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	var p Player
	err = query.QueryRow(username).Scan(&p.Username, &p.Balance, &p.FakeBalance, &p.ItemCount)
	if err != nil {
		return nil, fmt.Errorf("error scanning player: %v", err)
	}
	p.Balance = crypto_utils.RoundToTwoDecimals(p.Balance)
	p.FakeBalance = crypto_utils.RoundToTwoDecimals(p.FakeBalance)
	return &p, nil
}
