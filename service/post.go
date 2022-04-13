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

func (r *PostService) CreatePost(ctx context.Context, req *pp.CreatePostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().CreatePost(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) GetPost(ctx context.Context, req *pp.GetPostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().GetPost(req.PostId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) UpdatePost(ctx context.Context, req *pp.UpdatePostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().UpdatePost(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) DeletePost(ctx context.Context, req *pp.DeletePostReq) (*pp.Message, error) {
	res, err := r.storage.Post().DeletePost(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) Like(ctx context.Context, req *pp.LikePostReq) (*pp.Empty, error) {
	res, err := r.storage.Post().Like(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *PostService) ListUserPosts(ctx context.Context, req *pp.ListPostsReq) (*pp.ListPostsRes, error) {
	res, err := r.storage.Post().UserPostsList(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
