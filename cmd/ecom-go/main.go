package main

import (
	"database/sql"
	"log"
	"strconv"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/AntonioMorales97/ecom-go/internal/config"
	"github.com/AntonioMorales97/ecom-go/internal/logger"
	"github.com/AntonioMorales97/ecom-go/pkg/api"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/ecom?sslmode=disable"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to init Viper", zap.Error(err))
	}

	err = logger.InitializeZapCustomLogger(config.Env)
	if err != nil {
		log.Fatal("failed to init logger")
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Server.Host + ":" + strconv.Itoa(config.Server.Port))
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}
