package sqlite

var (
	createTable = `CREATE TABLE IF NOT EXISTS users (user_id INTEGER PRIMARY KEY, user_name TEXT UNIQUE ON CONFLICT IGNORE);
	CREATE TABLE IF NOT EXISTS reminds (remind_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), message TEXT, date INTEGER, period INTEGER, reminded BOOLEAN);
	CREATE TABLE IF NOT EXISTS urls (url_id INTEGER PRIMARY KEY, user_id REFERENCES users (user_id), url TEXT);`
	insertNewUser = `INSER INTO users (user_name) VALUES (?)`
	insertURL = `INSERT INTO urls (user_id, url) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?)`
	insertRemind = `INSERT INTO reminds (user_id, message, date, period, reminded) VALUES ((SELECT user_id FROM users WHERE user_name = ?), ?, ?, ?, 0)`
	pickRandom = `SELECT url FROM urls WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?) ORDER BY RANDOM() LIMIT 1`
	deleteUser = `DELETE FROM users WHERE user_name = ?`
	deleteURL = `DELETE FROM urls WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?)`
	deleteRemind = `DELETE FROM reminds WHERE user_id = (SELECT user_id FROM users WHERE user_name = ?)`
)
