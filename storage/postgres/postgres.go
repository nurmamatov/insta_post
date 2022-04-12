package postgres

import (
	"github.com/jmoiron/sqlx"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) CreatePost(req *pp.CreatePostReq) (*pp.CreatePostResp,error) {
	return nil,nil
}
func (r *PostRepo) GetPost(req *pp.GetPostReq) (*pp.GetPostRes,error) {
	return nil,nil
}
func (r *PostRepo) UpdatePost(req *pp.UpdatePostReq) (*pp.GetPostRes,error) {
	return nil,nil
}
func (r *PostRepo) DeletePost(req *pp.DeletePostReq) (*pp.Message,error) {
	return nil,nil
}