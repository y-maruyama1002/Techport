package repository

import (
	"github.com/y-maruyama1002/Techport/domain/entity"
)

type IBlogRepository interface {
	GetBlog(blogId string) (*entity.Blog, error)
	CreateBlog(blog *entity.Blog) error
}
