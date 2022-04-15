package postgres

import (
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	"time"
)

func (r *PostRepo) UpdatePost(req *pp.UpdatePostReq) (*pp.GetPostRes, error) {

	res, err := r.GetPost(req.PostId)
	if err != nil && res.PostId == "" {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339)
	queryPost := `UPDATE post SET title=$2, description=$3, updated_at=$4 WHERE post_id=$1 AND deleted_at IS NULL RETURNING post_id`
	_, err = r.db.Exec(queryPost, req.PostId, req.Title, req.Description, now)
	if err != nil {
		log.Println("Error while Update post", err)
		return nil, err
	}
	return r.GetPost(req.PostId)
}
