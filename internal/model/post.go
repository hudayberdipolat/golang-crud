package model

import "time"

type Post struct {
	ID              int       `json:"id"`
	PostTitle       string    `json:"post_title"`
	PostSlug        string    `json:"post_slug"`
	PostDescription string    `json:"post_description"`
	PostImageURLs   string    `json:"post_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
