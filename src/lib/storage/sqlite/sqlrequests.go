package storage

var (
	// Creator
	createTables = `CREATE TABLE IF NOT EXISTS users (user_id INTEGER PRIMARY KEY, user_name TEXT UNIQUE ON CONFLICT IGNORE);
	CREATE TABLE IF NOT EXISTS reminds (remind_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), message TEXT, date INTEGER, period INTEGER, reminded BOOLEAN);
	CREATE TABLE IF NOT EXISTS urls (url_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), url TEXT)`

	// Inserts
	insertNewUser = `INSERT INTO users (user_name) VALUES (?)`
	insertURL = `INSERT INTO urls (user_id, url) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?)`
	insertRemind = `INSERT INTO reminds (user_id, message, date, period, reminded) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?, ?, ?, 0)`

	// Randomize
	pickRandom = `SELECT url FROM urls WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?) ORDER BY RANDOM() LIMIT 1`

	// Delete
	deleteUser = `DELETE FROM users WHERE user_name = ?`
	deleteURL = `DELETE FROM urls WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?)`
	deleteRemind = `DELETE FROM reminds WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?)`

	// IsExists
	isExistsUser = `SELECT COUNT(*) FROM users WHERE user_name = ?`
	isExistsURL = `SELECT COUNT(*) FROM urls WHERE url = ? AND user_id = (SELECT user_id FROM users WHERE user_name = ?)`
	// For reminds mb need to check the message ???
	isExistsRemind = `SELECT COUNT(*) FROM reminds WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?) and date = ? and period = ?`
)
