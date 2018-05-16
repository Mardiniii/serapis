package database

// Statements for users table
const usersTable = `CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  api_key TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);`

const createUser = `INSERT INTO users (email, api_key)
VALUES ($1, $2)
RETURNING id, created_at`

const deleteUser = `DELETE FROM users
WHERE id = $1;`

const userByEmail = `SELECT * FROM users WHERE email=$1;`

const allUsers = `SELECT * FROM users`

// Staments for evaluations table
const evaluationsTable = `CREATE TABLE IF NOT EXISTS evaluations (
  id serial PRIMARY KEY,
  user_id INT NOT NULL,
  status VARCHAR(255) NOT NULL,
  language VARCHAR(255) NOT NULL,
	code TEXT,
	stdin TEXT[] DEFAULT '{}'::text[],
	dependencies JSONB,
	git VARCHAR(255),
	output TEXT,
	exit_code INT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);`

const createEvaluation = `INSERT INTO evaluations (user_id, status, language, code, stdin, dependencies, git)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at`

const evaluationByID = `SELECT * FROM evaluations WHERE id=$1;`
