package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/note-taking/api"
	_ "github.com/SaidovZohid/note-taking/api/docs"
	"github.com/SaidovZohid/note-taking/config"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/go-redis/redis/v9"
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

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	strg := storage.NewStorage(psqlConn)
	inMemory := storage.NewRedisStorage(rdb)

	server := api.New(&api.RouteOptions{
		Cfg: &cfg,
		Storage: strg,
		InMemory: inMemory,
	})

	err = server.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("Server Stoppen unexpectadly: %v", err)
	}
}
