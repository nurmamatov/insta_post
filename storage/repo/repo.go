package repo

import (
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

type PostStorageI interface {
	CreatePost(*pp.CreatePostReq) (*pp.CreatePostResp, error)
	GetPost(*pp.GetPostReq) (*pp.GetPostRes, error)
	// ListPost(*pp.ListPostReq) (*pp.PostResp, error)
	UpdatePost(*pp.UpdatePostReq) (*pp.GetPostRes, error)
	DeletePost(*pp.DeletePostReq) (*pp.Message, error)
	// UserPostsList(req *pp.CreatePost) (*pp.PostResp, error)
}
