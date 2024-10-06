package utils

type DbTransaction interface {
	Commit() error
	Rollback() error
}
