CREATE TABLE IF NOT EXISTS expense (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	description TEXT,
	amount INTEGER,
	category INTEGER,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

