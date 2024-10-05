package repositories

import "fmt"

type RepositoryError struct {
	Entity  string
	Message string
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Entity, e.Message)
}
