package postgres

import (
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) UpdatePost(req *pp.UpdatePostReq) (*pp.GetPostRes, error) {
	
	return nil, nil
}
func (r *PostRepo) DeletePost(req *pp.DeletePostReq) (*pp.Message, error) {
	return nil, nil
}
func (r *PostRepo) UserPostsList(req *pp.ListPostsReq) (*pp.ListPostsRes, error) {
	return nil, nil
}
func (r *PostRepo) Like(req *pp.LikePostReq) (*pp.Empty, error) {
	return nil, nil
}
