package domain

import "time"

type Blog struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBlog struct {
	Title string `json:"title" form:"title"`
	Body string `json:"body" form:"body"`
}

type BlogRepository interface {
	GetById(id int64) (Blog, error)
	CreateBlog(blog *CreateBlog) error
}

type BlogUsecase interface {
	GetById(id int64) (Blog, error)
	CreateBlog(blog *CreateBlog) error
}
