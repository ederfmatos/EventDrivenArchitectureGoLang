package repository

type UnitOfWork interface {
	Register(name string, repository interface{})
	GetRepository(name string) (interface{}, error)
	Do(fn func() error) error
}
