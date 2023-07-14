package main

import (
	"stats-pg/config"
	"stats-pg/repository/postgres"
	"stats-pg/server"
)

func main() {

	configParams := config.InitConfig("config")
	postgresRepo, err := postgres.NewPostgresRepository(configParams.GetString("database.connection.integration"))
	if err != nil {
		panic(err)
	}
	defer postgresRepo.Close()

	httpServer := server.InitHttpServer(postgresRepo)
	httpServer.Start()

}
