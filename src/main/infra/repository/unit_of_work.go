package repository

import (
	"EventDrivenArchitectureGoLang/src/main/application/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type DefaultUnitOfWork struct {
	DataBase     *sql.DB
	Tx           *sql.Tx
	Repositories map[string]repository.Factory
}

func NewDefaultUnitOfWork(dataBase *sql.DB) *DefaultUnitOfWork {
	return &DefaultUnitOfWork{
		DataBase:     dataBase,
		Repositories: make(map[string]repository.Factory),
	}
}

func (unitOfWork *DefaultUnitOfWork) Register(name string, repositoryFactory repository.Factory) {
	unitOfWork.Repositories[name] = repositoryFactory
}

func (unitOfWork *DefaultUnitOfWork) UnRegister(name string) {
	delete(unitOfWork.Repositories, name)
}

func (unitOfWork *DefaultUnitOfWork) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if unitOfWork.Tx == nil {
		tx, err := unitOfWork.DataBase.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		unitOfWork.Tx = tx
	}
	repo := unitOfWork.Repositories[name](unitOfWork.Tx)
	return repo, nil
}

func (unitOfWork *DefaultUnitOfWork) Do(ctx context.Context, fn func(Uow *DefaultUnitOfWork) error) error {
	if unitOfWork.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := unitOfWork.DataBase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	unitOfWork.Tx = tx
	err = fn(unitOfWork)
	if err != nil {
		errRb := unitOfWork.Rollback()
		if errRb == nil {
			return err
		}
		return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
	}
	return unitOfWork.CommitOrRollback()
}

func (unitOfWork *DefaultUnitOfWork) Rollback() error {
	if unitOfWork.Tx == nil {
		return errors.New("no transaction to rollback")
	}
	err := unitOfWork.Tx.Rollback()
	if err != nil {
		return err
	}
	unitOfWork.Tx = nil
	return nil
}

func (unitOfWork *DefaultUnitOfWork) CommitOrRollback() error {
	err := unitOfWork.Tx.Commit()
	if err == nil {
		unitOfWork.Tx = nil
		return nil
	}
	errRb := unitOfWork.Rollback()
	if errRb != nil {
		return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
	}
	return err
}
