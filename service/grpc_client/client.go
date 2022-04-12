package grpcclient

import (
	"fmt"
	"tasks/Instagram_clone/insta_post/config"
	pc "tasks/Instagram_clone/insta_post/genproto/comment_proto"
	pu "tasks/Instagram_clone/insta_post/genproto/user_proto"

	"google.golang.org/grpc"
)

type GrpcClientI interface {
	CommentService() pc.CommentServiceClient
	UserService() pu.UserServiceClient
}

// Client
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to comment service: host:%s port:%d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to user service: host:%s port:%d", cfg.UserServiceHost, cfg.UserServicePort)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"comment_service": pc.NewCommentServiceClient(connComment),
			"user_service":    pu.NewUserServiceClient(connUser),
		},
	}, nil
}

func (g *GrpcClient) CommentService() pc.CommentServiceClient {
	return g.connections["comment_service"].(pc.CommentServiceClient)
}

func (g *GrpcClient) UserService() pu.UserServiceClient {
	return g.connections["user_service"].(pu.UserServiceClient)
}
