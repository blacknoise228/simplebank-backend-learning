package gapi

import (
	"fmt"

	db "github.com/blacknoise228/simplebank-backend-learning/db/sqlc"
	"github.com/blacknoise228/simplebank-backend-learning/pb"
	"github.com/blacknoise228/simplebank-backend-learning/token"
	"github.com/blacknoise228/simplebank-backend-learning/util"
)

// Server server gRPC requests for out banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
