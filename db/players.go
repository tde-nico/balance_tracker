package db

import "fmt"

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
		if err != nil {
			return nil, fmt.Errorf("error scanning players: %v", err)
		}
		players = append(players, p)
	}
	return players, nil
}
