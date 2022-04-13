package main

import (
	"fmt"
	"log"
	"net"
	"tasks/Instagram_clone/insta_post/config"
	_ "github.com/lib/pq"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	"tasks/Instagram_clone/insta_post/service"
	grpcClient "tasks/Instagram_clone/insta_post/service/grpc_client"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.Load()

	psqlText := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDatabase,
	)

	grpcClient, err := grpcClient.New(config)
	if err != nil {
		log.Fatal("grpc dial error", err)
	}

	psql, err := sqlx.Connect("postgres", psqlText)
	if err != nil {
		log.Fatal(err)
	}

	PostService := service.NewPostService(psql, grpcClient)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal("Error while listening:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pp.RegisterPostServiceServer(s, PostService)
	log.Println("Main server runnning", config.Port)

	if err = s.Serve(lis); err != nil {
		log.Fatal("Error while listening:", err)
	}
}
