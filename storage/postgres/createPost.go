package postgres

import (
	"log"
	"strings"
	pp "tasks/Instagram_clone/insta_post/genproto/post_proto"
	"time"

	"github.com/google/uuid"
)

func (r *PostRepo) CreatePost(req *pp.CreatePostReq) (*pp.GetPostRes, error) {
	var (
		img    []string
		imgId  string
		postId string
	)
	queryPhoto := `
			INSERT INTO post_photo(
				image_id, 
				type,
				base_code,
				created_at
			)
			VALUES($1,$2,$3,$4)
			RETURNING image_id`
	queryPost := `
			INSERT INTO post(
				post_id,
				user_id,
				title,
				description,
				image,
				created_at
			)
			VALUES ($1,$2,$3,$4,$5,$6)
			RETURNING post_id`
	now := time.Now().Format(time.RFC3339)

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Error while begin tx", err)
		return nil, err
	}
	img = strings.Split(req.Image, ",")

	err = tx.QueryRow(queryPhoto, uuid.New(), img[0], img[1], now).Scan(
		&imgId,
	)
	if err != nil {
		log.Println("Error while insert post_photo", err)
		tx.Rollback()
		return nil, err
	}

	err = tx.QueryRow(queryPost, uuid.New(), req.UserId, req.Title, req.Description, imgId, now).Scan(
		&postId,
	)
	if err != nil {
		log.Println("Error while insert post", err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return r.GetPost(postId)
}
