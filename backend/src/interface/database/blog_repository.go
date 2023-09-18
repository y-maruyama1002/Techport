package database

import (
	"github.com/y-maruyama1002/Techport/domain/entity"
	"github.com/y-maruyama1002/Techport/domain/repository"
	"gorm.io/gorm"
)

type BlogRepository struct {
	*gorm.DB
}

func NewBlogRepository(db *gorm.DB) repository.IBlogRepository {
	return &BlogRepository{db}
}

func (r *BlogRepository) GetBlog(blogId string) (*entity.Blog, error) {
	blog := entity.Blog{}
	r.First(&blog, blogId)
	return &blog, nil
}
