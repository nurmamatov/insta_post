package service

import (
	"context"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	grpcClient "tasks/Instagram_clone/insta_post/service/grpc_client"
	"tasks/Instagram_clone/insta_post/storage"

	"github.com/jmoiron/sqlx"
)

type PostService struct {
	storage storage.IStorage
	client  grpcClient.GrpcClientI
}

func NewPostService(db *sqlx.DB, client grpcClient.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		client:  client,
	}
}

func (r *PostService) CreatePost(ctx context.Context, req *pp.CreatePostReq) (*pp.CreatePostResp, error) {
	return nil, nil
}
func (r *PostService) GetPost(ctx context.Context, req *pp.GetPostReq) (*pp.GetPostRes, error) {
	return nil, nil
}
func (r *PostService) UpdatePost(ctx context.Context, req *pp.UpdatePostReq) (*pp.GetPostRes, error) {
	return nil, nil
}
func (r *PostService) DeletePost(ctx context.Context, req *pp.DeletePostReq) (*pp.Message, error) {
	return nil, nil
}
