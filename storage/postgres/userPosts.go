package postgres

import (
	"fmt"
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

func (r *PostRepo) UserPostsList(req *pp.ListPostsReq) (*pp.ListPostsRes, error) {

	var (
		resList pp.ListPostsRes
	)
	queryUser := `SELECT post_id FROM post WHERE user_id=$1 AND deleted_at IS NULL`
	rows, err := r.db.Query(queryUser, req.UserId)
	if err != nil {
		log.Println("Error while Users posts list", err)
		return nil, err
	}
	
	for rows.Next() {
		postId := ""
		err = rows.Scan(
			&postId,
		)
		if err != nil {
			log.Println("Error while take postsId with userId", err)
			return nil, err
		}
		res, err := r.GetPost(postId)
		if err != nil {
			fmt.Println("Error while get user in list method", err)
			return nil, err
		}
		resList.Posts = append(resList.Posts, res)
	}
	
	return &resList, nil
}
