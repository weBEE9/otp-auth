package model

const Schema = `CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	phone_number TEXT NOT NULL UNIQUE,
	phone_verified BOOLEAN NOT NULL DEFAULT FALSE
);`