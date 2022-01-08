package api

import (
	"os"
	"testing"
	"time"

	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	token := util.TokenConfig{
		SymmetricKey:   util.RandomString(32),
		AccessDuration: time.Minute,
	}
	config := util.Config{
		Token: token,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
