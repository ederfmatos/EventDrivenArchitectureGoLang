package repository

type SqlConnection interface {
	Update(command string, args ...any) error
}
