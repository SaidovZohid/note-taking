package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/note-taking/api"
	"github.com/SaidovZohid/note-taking/config"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/SaidovZohid/note-taking/api/docs"
)

func main() {
	cfg := config.New(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("Configuration: ", cfg)
	fmt.Println("Connected Succesfully!")

	strg := storage.New(psqlConn)

	server := api.New(&api.RouteOptions{
		Cfg: &cfg,
		Storage: &strg,
	})

	err = server.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("Server Stoppen unexpectadly: %v", err)
	}
}
