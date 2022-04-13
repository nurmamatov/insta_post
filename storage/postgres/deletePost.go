package postgres

import (
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	"time"
)

func (r *PostRepo) DeletePost(req *pp.DeletePostReq) (*pp.Message, error) {

	now := time.Now().Format(time.RFC3339)
	queryPost := `UPDATE post SET deleted_at=$2 WHERE post_id=$1`

	res, err := r.GetPost(req.PostId)
	if err != nil && res.PostId == "" {
		return &pp.Message{Message: "This post haven't"}, nil
	}

	_, err = r.db.Exec(queryPost, req.PostId, now)
	if err != nil {
		log.Println("Error while Delete post", err)
		return nil, err
	}

	return &pp.Message{Message: "Deleted!"}, nil
}
