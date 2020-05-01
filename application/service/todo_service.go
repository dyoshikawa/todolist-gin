package service

type TodoService interface{}

type TodoServiceImpl struct{}

type TodoServiceParams struct {
	Body string
}

func (s *TodoServiceImpl) store(params TodoServiceParams) {

}
