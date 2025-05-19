package migrations

import "database/sql"

type createTodoTable struct{}

func (m *createTodoTable) SkipProd() bool {
	return false
}

func getCreateTodoTable() migration {
	return &createTodoTable{}
}

func (m *createTodoTable) Name() string {
	return "create-todo"
}

func (m *createTodoTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE todo (
			id SERIAL PRIMARY KEY,
			todo VARCHAR(255) NOT NULL,
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)	

	if err != nil {
		return err
	}
	return err
}

func (m *createTodoTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE todo")
	return err
}