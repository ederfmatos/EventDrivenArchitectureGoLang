package repository

import (
	"context"
	"database/sql"
)

type Factory func(tx *sql.Tx) interface{}

type UnitOfWork interface {
	Register(name string, repositoryFactory Factory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(unitOfWork *UnitOfWork) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}
