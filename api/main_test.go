package api

import (
	"os"
	"testing"

	db "github.com/blacknoise228/simplebank-backend-learning/db/sqlc"
	"github.com/gin-gonic/gin"
)

func newTestServer(store db.Store) *Server {

	return NewServer(store)
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
