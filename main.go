package main

import (
	"codechallenge/cmd"
	"embed"
)

var (
	//go:embed db/migrations
	MigrationFS embed.FS
)

//go:generate mockgen -source=internal/repository/repository.go -destination=mocks/repository_mock.go -package=mocks StorageRepository,QueueRepository,TodoRepository
//go:generate mockgen -source=utils/database.go -destination=mocks/utils_mock.go -package=mocks DbTransaction
//go:generate sqlboiler --wipe --add-soft-deletes psql -o internal/repository/database/boiler
func main() {
	cmd.Execute(MigrationFS)
}
