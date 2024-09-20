package to_do_service

import (
	"todo_app/entities"
	"todo_app/repositories/to_do_repository"
)

type ITODOService interface {
	Create(todoToCreate *entities.TODO) error
	List() ([]*entities.TODO, error)
	Get(id uint64) (*entities.TODO, error)
	Update(todoToUpdate *entities.TODO) error
	Delete(id uint64) error
}

type TODOService struct {
	todoRepository *to_do_repository.TODORepository

	ITODOService
}

func New(todoRepository *to_do_repository.TODORepository) *TODOService {
	return &TODOService{
		todoRepository: todoRepository,
	}
}

func (dts *TODOService) Create(todoToCreate *entities.TODO) error {
	return dts.todoRepository.Create(todoToCreate)
}

func (dts *TODOService) List() ([]*entities.TODO, error) {
	return dts.todoRepository.List()
}

func (dts *TODOService) Get(id uint64) (*entities.TODO, error) {
	return dts.todoRepository.Get(id)
}

func (dts *TODOService) Update(todoToUpdate *entities.TODO) error {
	return dts.todoRepository.Update(todoToUpdate)
}

func (dts *TODOService) Delete(id uint64) error {
	return dts.todoRepository.Delete(id)
}
