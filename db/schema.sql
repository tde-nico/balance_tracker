CREATE TABLE IF NOT EXISTS "keys" (
	"name" TEXT NOT NULL,
	"key" TEXT UNIQUE NOT NULL,
	PRIMARY KEY("name")
);

CREATE TABLE IF NOT EXISTS "users" (
	"username" VARCHAR(32) NOT NULL,
	"salt" VARCHAR(16) NOT NULL,
	"password" VARCHAR(64) NOT NULL,
	"apikey" VARCHAR(256) NOT NULL,
	"is_editor" BOOLEAN NOT NULL DEFAULT 0,
	PRIMARY KEY("username")
);

CREATE TABLE IF NOT EXISTS "players" (
	"username" VARCHAR(32) NOT NULL,
	"balance" REAL NOT NULL DEFAULT 0,
	"fake_balance" REAL NOT NULL DEFAULT 0,
	"item_count" INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY("username")
);

CREATE TABLE IF NOT EXISTS "out_transactions" (
	"id" INTEGER NOT NULL,
	"from" VARCHAR(32) NOT NULL,
	"to" VARCHAR(32) NOT NULL,
	"purpose" TEXT NOT NULL,
	"amount" REAL NOT NULL,
	"is_arrived" BOOLEAN NOT NULL DEFAULT 0,
	"note" TEXT,
	PRIMARY KEY("id"),
	FOREIGN KEY("from") REFERENCES "players"("username"),
	FOREIGN KEY("to") REFERENCES "players"("username")
);

CREATE TABLE IF NOT EXISTS "peer_transactions" (
	"id" INTEGER NOT NULL,
	"from" VARCHAR(32) NOT NULL,
	"to" VARCHAR(32) NOT NULL,
	"amount" REAL NOT NULL,
	"is_arrived" BOOLEAN NOT NULL DEFAULT 0,
	"note" TEXT,
	PRIMARY KEY("id"),
	FOREIGN KEY("from") REFERENCES "players"("username"),
	FOREIGN KEY("to") REFERENCES "players"("username")
);

CREATE TABLE IF NOT EXISTS "in_transactions" (
	"id" INTEGER NOT NULL,
	"from" VARCHAR(255) NOT NULL,
	"to" VARCHAR(32) NOT NULL,
	"amount" REAL NOT NULL,
	"is_arrived" BOOLEAN NOT NULL DEFAULT 0,
	"note" TEXT,
	PRIMARY KEY("id"),
	FOREIGN KEY("to") REFERENCES "players"("username")
);
