package mysql_repository

import "github.com/go-gorp/gorp"

type TodoMysqlRepository struct {
	db *gorp.DbMap
}

type CreateTodoParams struct {
	body string
}

func (r *TodoMysqlRepository) createTodo(params CreateTodoParams) error {
	_, err := r.db.Exec("INSERT INTO todos(body) VALUES(?)")
	if err != nil {
		return err
	}
	return nil
}
