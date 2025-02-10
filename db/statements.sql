-- GetKey
SELECT key
	FROM keys
	WHERE name = ?;

-- UserExists
SELECT username
	FROM users
	WHERE username = ?;

-- GetUserByAPIKey
SELECT username, is_editor
	FROM users
	WHERE apikey = ?;

-- GetUserByUsername
SELECT username, is_editor
	FROM users
	WHERE username = ?;

-- CreateUser
INSERT INTO users (username, salt, password, apikey, is_editor)
	VALUES (?, ?, ?, ?, ?);

-- LoginUser
SELECT apikey, salt, password
	FROM users
	WHERE username = ?;

-- GetTotalBalance
SELECT SUM(balance), SUM(balance + fake_balance)
	FROM players;

-- GetNotedTransactions
SELECT note, amount
	FROM in_transactions
	WHERE note NOT NULL;

-- GetPlayers
SELECT username, balance, fake_balance, item_count
	FROM players;
