package main

import (
	"database/sql"
	"log"
	"strconv"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/AntonioMorales97/ecom-go/pkg/api"
	"github.com/AntonioMorales97/ecom-go/pkg/util"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to init Viper", zap.Error(err))
	}

	err = util.InitializeZapCustomLogger(config.Env)
	if err != nil {
		log.Fatal("failed to init logger")
	}

	conn, err := sql.Open(config.Db.Driver, config.Db.Source)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("failed to create server", err)
	}

	err = server.Start(config.Server.Host + ":" + strconv.Itoa(config.Server.Port))
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
