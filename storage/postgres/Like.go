package postgres

import (
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"

	"github.com/google/uuid"
)

func (r *PostRepo) Like(req *pp.LikePostReq) (bool, error) {

	query := `INSERT INTO likes (like_id, post_id, user_id, likes) VALUES($1,$2,$3,$4)`
	_, err := r.db.Exec(query, uuid.New(), req.PostId, req.UserId, true)
	if err != nil {
		log.Println("Error while insert like", err)
		return false, err
	}

	return true, nil
}
func (r *PostRepo) CheckLike(req *pp.GetPostReq) (bool, error) {
	var (
		checkLike bool
	)
	query := `SELECT likes FROM likes WHERE user_id=$1 AND post_id=$2`
	err := r.db.QueryRow(query, req.UserId, req.PostId).Scan(&checkLike)
	if err != nil {
		return false, err
	}
	if checkLike == bool(true) {
		return true, nil
	} else {
		return false, nil
	}
}
func (r *PostRepo) DeleteLike(req *pp.LikeDeleteReq) (*pp.Bool, error) {
	query := `DELETE FROM likes WHERE post_id=$1 AND user_id=$2`
	_, err := r.db.Exec(query, req.PostId, req.UserId)
	if err != nil {
		log.Println("Error while delete like", err)
		return &pp.Bool{Result: true}, err
	}
	return &pp.Bool{Result: false}, nil
}
