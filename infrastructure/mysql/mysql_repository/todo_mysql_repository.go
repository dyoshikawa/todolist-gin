package mysql_repository

type TodoMysqlRepository struct{}

type CreateTodoParams struct {
	body string
}

func createTodo(params CreateTodoParams) error {
	return nil
}
