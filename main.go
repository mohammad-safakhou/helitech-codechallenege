package main

import "codechallenge/cmd"

//go:generate mockgen -source=internal/repository/repository.go -destination=mocks/repository_mock.go -package=mocks StorageRepository,QueueRepository,TodoRepository
//go:generate mockgen -source=utils/database.go -destination=mocks/utils_mock.go -package=mocks DbTransaction
func main() {
	cmd.Execute()
}
