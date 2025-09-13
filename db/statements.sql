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
	FROM players
	ORDER BY username;

-- GetPlayer
SELECT username, balance, fake_balance, item_count
	FROM players
	WHERE username = ?;

-- GetInTransactions
SELECT id, "from", "to", amount, is_arrived, note
	FROM in_transactions
	WHERE "to" = ?
	ORDER BY id DESC;

-- GetAllInTransactions
SELECT id, "from", "to", amount, is_arrived, note
	FROM in_transactions
	ORDER BY id DESC;

-- GetInTransaction
SELECT id, "from", "to", amount, is_arrived, note
	FROM in_transactions
	WHERE id = ?;

-- InsertInTransaction
INSERT INTO in_transactions ("from", "to", amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?);

-- InsertInTransactionWithID
INSERT INTO in_transactions (id, "from", "to", amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?, ?);

-- DeleteInTransaction
DELETE FROM in_transactions
	WHERE id = ?;

-- GetOutTransactions
SELECT id, "from", "to", purpose, amount, is_arrived, note
	FROM out_transactions
	WHERE "from" = ?
		OR "to" = ?
	ORDER BY id DESC;

-- GetAllOutTransactions
SELECT id, "from", "to", purpose, amount, is_arrived, note
	FROM out_transactions
	ORDER BY id DESC;

-- GetOutTransaction
SELECT id, "from", "to", purpose, amount, is_arrived, note
	FROM out_transactions
	WHERE id = ?;

-- InsertOutTransaction
INSERT INTO out_transactions ("from", "to", purpose, amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?, ?);

-- InsertOutTransactionWithID
INSERT INTO out_transactions (id, "from", "to", purpose, amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?, ?, ?);

-- DeleteOutTransaction
DELETE FROM out_transactions
	WHERE id = ?;

-- GetPeerTransactions
SELECT id, "from", "to", amount, is_arrived, note
	FROM peer_transactions
	WHERE "from" = ?
		OR "to" = ?
	ORDER BY id DESC;

-- GetAllPeerTransactions
SELECT id, "from", "to", amount, is_arrived, note
	FROM peer_transactions
	ORDER BY id DESC;

-- GetPeerTransaction
SELECT id, "from", "to", amount, is_arrived, note
	FROM peer_transactions
	WHERE id = ?;

-- InsertPeerTransaction
INSERT INTO peer_transactions ("from", "to", amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?);

-- InsertPeerTransactionWithID
INSERT INTO peer_transactions (id, "from", "to", amount, is_arrived, note)
	VALUES (?, ?, ?, ?, ?, ?);

-- DeletePeerTransaction
DELETE FROM peer_transactions
	WHERE id = ?;
