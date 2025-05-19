package migrations

import "database/sql"

type createUserTable struct{}

func (m *createUserTable) SkipProd() bool {
	return false
}

func getCreateUserTable() migration {
	return &createUserTable{}
}

func (m *createUserTable) Name() string {
	return "create-user"
}

func (m *createUserTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)

	if err != nil {
		return err
	}

	return err
}

func (m *createUserTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE users")

	return err
}
