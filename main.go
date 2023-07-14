package main

import (
	"stats-pg/repository/postgres"
	"stats-pg/server"
)

func main() {

	postgresRepo, err := postgres.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer postgresRepo.Close()

	httpServer := server.InitHttpServer(postgresRepo)
	httpServer.Start()

}
