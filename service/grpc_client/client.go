package grpcclient

import (
	"fmt"
	"tasks/Instagram_clone/insta_post/config"
	pc "tasks/Instagram_clone/insta_post/genproto/comment_proto"

	"google.golang.org/grpc"
)

// GrpcClientI ...
type GrpcClientI interface {
	CommentService() pc.CommentServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("comment service dial host:%s port:%d err:%s",
			cfg.CommentServiceHost,
			cfg.CommentServicePort,
			err.Error())
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"comment_service": pc.NewCommentServiceClient(connComment),
		},
	}, nil
}

func (g *GrpcClient) CommentService() pc.CommentServiceClient {
	return g.connections["comment_service"].(pc.CommentServiceClient)
}
