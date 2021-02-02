package database

// CreateItemTable ...
func CreateItemTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name varchar(100),
		date timestamp default current_timestamp
	)
	`)
}

// CreateUserTable ...
func CreateUserTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name varchar(100),
		email varchar(30) UNIQUE,
		password varchar(100),
		register_date timestamp default current_timestamp
	)
	`)
}