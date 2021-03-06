package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("failed to init Viper", zap.Error(err))
	}

	err = util.InitializeZapCustomLogger(config.Env)
	if err != nil {
		log.Fatal("failed to init logger")
	}

	testDB, err = sql.Open(config.Db.Driver, config.Db.Source)
	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
