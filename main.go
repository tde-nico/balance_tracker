package main

import (
	"balance/db"
	"balance/log"
	"balance/routes"
	"encoding/hex"
	"flag"
	"os"
	"strings"
)

type Flags struct {
	migrate  bool
	triggers bool
	register string
	prune    bool
	dev      bool
}

func InitFlags() *Flags {
	var flags Flags
	flag.BoolVar(&flags.migrate, "m", false, "Migrate the database")
	flag.BoolVar(&flags.triggers, "t", false, "Reload Triggers")
	flag.StringVar(&flags.register, "r", "", "Register a new user")
	flag.BoolVar(&flags.prune, "p", false, "Clean ALL the database")
	flag.BoolVar(&flags.dev, "d", false, "Enables Dev Mode")
	flag.Parse()
	if _, err := os.Stat("DEV"); err == nil || os.Getenv("DEV") != "" {
		flags.dev = true
	}
	return &flags
}

func EvalFlags(flags *Flags) bool {
	if flags.dev {
		log.Notice("Dev mode enabled")
	}

	if flags.migrate {
		db.DropTables()
		db.ExecSQLFile("db/schema.sql")
		db.ExecSQLFile("db/triggers.sql")

	} else if flags.triggers {
		db.ExecSQLFile("db/triggers.sql")

	} else if flags.register != "" {
		data := strings.Split(flags.register, ":")
		if len(data) != 3 {
			log.Fatal("Invalid register format")
		}
		user := data[0]
		password := data[1]
		isEditor := data[2] == "1"
		log.Noticef("Registering user \"%s\" with password \"%s\" and isEditor %v", user, password, isEditor)

		err := db.RegisterUser(user, password, isEditor)
		if err != nil {
			log.Fatalf("Error registering user: %v", err)
		}

		log.Noticef("User Successfully registered")

	} else if flags.prune {
		db.PruneDB()

	} else {
		return true
	}
	return false
}

func LoadSecretKey() []byte {
	var err error

	secretHex, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Info("SECRET_KEY not found in environment, using database")
		secretHex, err = db.GetKey("SECRET_KEY")
		if err != nil {
			log.Fatalf("SECRET_KEY not found: %v", err)
		}
	}

	if len(secretHex) != 64 {
		log.Fatal("SECRET_KEY must be 32 bytes long")
	}

	secret := make([]byte, 32)
	n, err := hex.Decode(secret, []byte(secretHex))
	if err != nil {
		log.Fatalf("Error decoding SECRET_KEY: %v", err)
	}
	if n != 32 {
		log.Fatal("Error decoding SECRET_KEY")
	}

	return secret
}

func main() {
	flags := InitFlags()

	db.InitDB("database.db")
	defer db.CloseDB()
	db.LoadStatements("db/statements.sql")
	if !EvalFlags(flags) {
		return
	}

	key := LoadSecretKey()
	routes.StartRouting(key)
}
