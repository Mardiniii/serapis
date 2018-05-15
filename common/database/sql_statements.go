package database

const usersTable = `CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  api_key TEXT UNIQUE NOT NULL,
  created_at timestamp with time zone DEFAULT current_timestamp
);`

const createUser = `INSERT INTO users (email, api_key)
VALUES ($1, $2)
RETURNING id`

const evaluationsTable = `CREATE TABLE IF NOT EXISTS evaluations (
  id serial PRIMARY KEY,
  user_id INT NOT NULL,
  status VARCHAR(255) NOT NULL,
  language VARCHAR(255) NOT NULL,
	code TEXT NOT NULL,
	stdin TEXT NOT NULL,
	dependencies TEXT NOT NULL,
	git VARCHAR(255) NOT NULL,
	output TEXT NOT NULL,
	exit_code INT NOT NULL,
  created_at timestamp with time zone DEFAULT current_timestamp
);`
