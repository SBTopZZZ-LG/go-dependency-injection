package to_do_repository

import (
	"github.com/a631807682/zerofield"
	"gorm.io/gorm"
	"todo_app/entities"
)

type ITODORepository interface {
	Create(todoToCreate *entities.TODO) error
	List() ([]*entities.TODO, error)
	Get(id uint64) (*entities.TODO, error)
	Update(todoToUpdate *entities.TODO) error
	Delete(id uint64) error
}

type TODORepository struct {
	gormDB *gorm.DB

	ITODORepository
}

func New(gormDB *gorm.DB) *TODORepository {
	return &TODORepository{
		gormDB: gormDB,
	}
}

func (tr *TODORepository) Create(todoToCreate *entities.TODO) error {
	result := tr.gormDB.Create(&todoToCreate).First(&todoToCreate)

	return result.Error
}

func (tr *TODORepository) List() ([]*entities.TODO, error) {
	var todos []*entities.TODO
	result := tr.gormDB.Order("is_important DESC").Find(&todos)

	return todos, result.Error
}

func (tr *TODORepository) Get(id uint64) (*entities.TODO, error) {
	var todo *entities.TODO
	result := tr.gormDB.First(&todo, id)

	return todo, result.Error
}

func (tr *TODORepository) Update(todoToUpdate *entities.TODO) error {
	result := tr.gormDB.Scopes(zerofield.UpdateScopes()).Updates(&todoToUpdate).First(&todoToUpdate)

	return result.Error
}

func (tr *TODORepository) Delete(id uint64) error {
	result := tr.gormDB.Delete(&entities.TODO{}, id)

	return result.Error
}
