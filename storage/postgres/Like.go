package postgres

import (
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

func (r *PostRepo) Like(req *pp.LikePostReq) (bool, error) {
	res, _ := r.CheckLike(req)
	if res == true {
		return false, nil
	}

	query := `INSERT INTO likes (post_id, user_id) VALUES($1,$2)`
	_, err := r.db.Exec(query, req.PostId, req.UserId)
	if err != nil {
		log.Println("Error while insert like", err)
		return false, err
	}

	return true, nil
}
func (r *PostRepo) CheckLike(req *pp.LikePostReq) (bool, error) {
	var (
		checkLike int
	)
	query := `SELECT COUNT(*) FROM likes WHERE user_id=$1 AND post_id=$2`
	err := r.db.QueryRow(query, req.UserId, req.PostId).Scan(&checkLike)
	if err != nil {
		return false, err
	}
	if checkLike == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
func (r *PostRepo) PostLikes(req *pp.LikePostReq) (int64, error) {
	var (
		countLike int64
	)
	query := `SELECT COUNT(*) FROM likes WHERE post_id=$1`
	err := r.db.QueryRow(query, req.PostId).Scan(&countLike)
	if err != nil {
		return 0, err
	}
	return countLike, nil
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
