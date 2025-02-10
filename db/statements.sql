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
SELECT balance - SUM(amount), fake - SUM(amount)
	FROM (
		SELECT SUM(balance) AS balance, SUM(balance + fake_balance) AS fake
			FROM players),
		in_transactions
	WHERE note NOT NULL;
