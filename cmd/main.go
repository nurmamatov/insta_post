package main

import (
	"fmt"
	"log"
	"net"
	"tasks/Instagram_clone/insta_post/config"
	pc "tasks/Instagram_clone/insta_post/genproto/comment_proto"
	pu "tasks/Instagram_clone/insta_post/genproto/post_proto"
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

	psql, err := sqlx.Connect("postgres", psqlText)
	if err != nil {
		log.Fatal(err)
	}

	client, err := grpcClient.New(config)
	if err != nil {
		log.Fatal("grpc dial error", err)
	}

	PostService := service.NewPostService(psql, client)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal("Error while listening: %v", err)
	}

	s := grpc.NewServer()
	pc.RegisterCommentServiceServer(s, PostService)
	pu.RegisterPostServiceServer(s, PostService)
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", err)
	}
}
