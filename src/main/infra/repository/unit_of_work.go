package repository

import (
	"gorm.io/gorm"
)

type DefaultUnitOfWork struct {
	DB           *gorm.DB
	Repositories map[string]interface{}
}

func NewDefaultUnitOfWork(DB *gorm.DB) *DefaultUnitOfWork {
	return &DefaultUnitOfWork{
		DB:           DB,
		Repositories: make(map[string]interface{}),
	}
}

func (unitOfWork *DefaultUnitOfWork) Register(name string, repository interface{}) {
	unitOfWork.Repositories[name] = repository
}

func (unitOfWork *DefaultUnitOfWork) GetRepository(name string) (interface{}, error) {
	return unitOfWork.Repositories[name], nil
}

func (unitOfWork *DefaultUnitOfWork) Do(fn func() error) error {
	return unitOfWork.DB.Transaction(func(*gorm.DB) error {
		return fn()
	})
}
