package main

import (
	"net"
	"tasks/Instagram_clone/insta_post/config"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	"tasks/Instagram_clone/insta_post/pkg/db"
	"tasks/Instagram_clone/insta_post/pkg/logger"
	"tasks/Instagram_clone/insta_post/service"
	grpcClient "tasks/Instagram_clone/insta_post/service/grpc_client"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post_service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postres error", logger.Error(err))
	}

	client, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("grpc dial error", logger.Error(err))
	}

	PostService := service.NewPostService(connDB, log, client)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pp.RegisterPostServiceServer(s, PostService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.Port))

	if err = s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
