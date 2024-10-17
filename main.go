package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/blacknoise228/simplebank-backend-learning/api"
	db "github.com/blacknoise228/simplebank-backend-learning/db/sqlc"
	"github.com/blacknoise228/simplebank-backend-learning/gapi"
	"github.com/blacknoise228/simplebank-backend-learning/pb"
	"github.com/blacknoise228/simplebank-backend-learning/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("DB not connected", err)
	}
	store := db.NewStore(conn)
	go runGateWayServer(config, store)
	runGRPCServer(config, store)

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("server create error:", err)
	}

	if err = server.Start(config.HTTPServerURL); err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runGRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("server create error:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerURL)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create gRPC server:", err)
	}
}
func runGateWayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("http gatewayserver create error:", err)
	}
	jsonOption := runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot create http gateway handler:", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerURL)
	if err != nil {
		log.Fatal("cannot create http gateway listener:", err)
	}
	log.Printf("start HTTP gateway server on %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot create http gateway server:", err)
	}
}
