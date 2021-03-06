package repo

import (
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

type PostStorageI interface {
	CreatePost(*pp.CreatePostReq) (*pp.GetPostRes, error)
	GetPost(string) (*pp.GetPostRes, error)
	UpdatePost(*pp.UpdatePostReq) (*pp.GetPostRes, error)
	DeletePost(*pp.DeletePostReq) (*pp.Message, error)
	UserPostsList(req *pp.ListPostsReq) (*pp.ListPostsRes, error)
	Like(req *pp.LikePostReq) (bool, error)
	DeleteLike(req *pp.LikeDeleteReq) (*pp.Bool, error)
}
