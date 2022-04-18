package postgres

import (
	"database/sql"
	"log"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
)

func (r *PostRepo) GetPost(postId string) (*pp.GetPostRes, error) {
	var (
		img_type  string
		base_code string
		countLike int64
	)
	res := pp.GetPostRes{}
	queryPost := `SELECT post_id, user_id, title, description, image, created_at FROM post WHERE post_id=$1 AND deleted_at IS NULL`
	queryLike := `SELECT COUNT(likes) FROM likes WHERE post_id=$1`
	queryPhoto := `SELECT type, base_code FROM post_photo WHERE image_id=$1 AND deleted_at IS NULL`

	err := r.db.QueryRow(queryPost, postId).Scan(
		&res.PostId,
		&res.UserId,
		&res.Title,
		&res.Description,
		&res.Image,
		&res.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error while get post:", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return &pp.GetPostRes{PostId: ""}, err
	}

	err = r.db.QueryRow(queryLike, postId).Scan(
		&countLike,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error while Get Like:", err)
		return nil, err
	}

	err = r.db.QueryRow(queryPhoto, res.Image).Scan(
		&img_type,
		&base_code,
	)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error while Get Photo:", err)
		return nil, err
	}
	res.Image = img_type + base_code
	res.Likes = countLike
	res.CheckLike, _ = r.CheckLike(&pp.LikePostReq{UserId: res.UserId, PostId: postId})
	res.Likes, _ = r.PostLikes(&pp.LikePostReq{PostId: postId})

	return &res, nil
}
