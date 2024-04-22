package storage

var (
	// Creator
	createTables = `CREATE TABLE IF NOT EXISTS users (user_id INTEGER PRIMARY KEY, user_name TEXT UNIQUE ON CONFLICT IGNORE, message_id INTEGER, chat_id INTEGER);
	CREATE TABLE IF NOT EXISTS reminds (remind_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), message TEXT, date INTEGER, period INTEGER, reminded BOOLEAN);
	CREATE TABLE IF NOT EXISTS urls (url_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), url TEXT, removed BOOLEAN);`

	// Inserts
	insertNewUser = `INSERT INTO users (user_name) VALUES (?)`
	insertURL = `INSERT INTO urls (user_id, url, removed) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?, FALSE)`
	insertRemind = `INSERT INTO reminds (user_id, message, date, period, reminded) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?, ?, ?, 0)`

	// Upadate
	updateUsersIDs = `UPDATE users SET message_id = ?, chat_id = ? WHERE user_name = ?`
	updateURLRemoved = `UPDATE urls SET removed = FALSE WHERE url = ? and user_id = (SELECT user_id FROM users WHERE user_name = ?)`

	// Randomize
	pickRandomURL = `SELECT url FROM urls WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?) AND removed = FALSE ORDER BY RANDOM() LIMIT 1`
	pickUserInfo = `SELECT message_id, chat_id FROM users WHERE user_name = ?`

	// Delete
	deleteUser = `DELETE FROM users WHERE user_name = ?`
	deleteURL = `UPDATE urls SET removed = TRUE WHERE url = ? AND user_id = (SELECT user_id FROM users WHERE user_name = ?)`
	// deleteRemind = `DELETE FROM reminds WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?)`

	// IsExists
	isExistsUser = `SELECT COUNT(*) FROM users WHERE user_name = ?`
	isExistsURL = `SELECT COUNT(*) FROM urls WHERE url = ? AND user_id = (SELECT user_id FROM users WHERE user_name = ? AND removed = FALSE)`
	isRemovedURL = `SELECT COUNT(*) FROM urls WHERE url = ? AND user_id = (SELECT user_id FROM users WHERE user_name = ? AND removed = TRUE)`
	// For reminds mb need to check the message ???
	// isExistsRemind = `SELECT COUNT(*) FROM reminds WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?) and date = ? and period = ?`
)
