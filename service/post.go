package service

import (
	"context"
	"fmt"
	pc "tasks/Instagram_clone/insta_post/genproto/comment_proto"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	l "tasks/Instagram_clone/insta_post/pkg/logger"

	grpcClient "tasks/Instagram_clone/insta_post/service/grpc_client"
	"tasks/Instagram_clone/insta_post/storage"

	"github.com/jmoiron/sqlx"
)

// PostServie ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (r *PostService) CreatePost(ctx context.Context, req *pp.CreatePostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().CreatePost(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *PostService) GetPost(ctx context.Context, req *pp.GetPostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().GetPost(req.PostId)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}

	ress, err := r.client.CommentService().GetComment(context.Background(), &pc.GetCommentReq{PostId: res.PostId})
	fmt.Println(ress)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	for _, j := range ress.Comments {
		res.Comments = append(res.Comments, &pp.Comment{UserId: j.UserId, Text: j.Text})
	}

	return res, nil
}
func (r *PostService) UpdatePost(ctx context.Context, req *pp.UpdatePostReq) (*pp.GetPostRes, error) {
	res, err := r.storage.Post().UpdatePost(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *PostService) DeletePost(ctx context.Context, req *pp.DeletePostReq) (*pp.Message, error) {
	res, err := r.storage.Post().DeletePost(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *PostService) Like(ctx context.Context, req *pp.LikePostReq) (*pp.Bool, error) {
	res, err := r.storage.Post().Like(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return &pp.Bool{Result: false}, err
	}
	return &pp.Bool{Result: res}, nil
}
func (r *PostService) DeleteLike(ctx context.Context, req *pp.LikeDeleteReq) (*pp.Bool, error) {
	res, err := r.storage.Post().DeleteLike(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return &pp.Bool{Result: true}, err
	}
	return res, nil
}
func (r *PostService) ListUserPosts(ctx context.Context, req *pp.ListPostsReq) (*pp.ListPostsRes, error) {
	res, err := r.storage.Post().UserPostsList(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
