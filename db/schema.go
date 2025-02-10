package db

type User struct {
	Username string
	ApiKey   string
	IsEditor bool
}

type Player struct {
	Username    string
	Balance     float32
	FakeBalance float32
	ItemCount   int
}

type Transaction struct {
	ID        int
	From      string
	To        string
	Purpose   string
	Amount    float32
	IsArrived bool
	Note      string
}
