package usecase

import (
	"github.com/y-maruyama1002/Techport/common/dto"
	"github.com/y-maruyama1002/Techport/domain/entity"
	"github.com/y-maruyama1002/Techport/domain/repository"
)

type BlogInputPort interface {
	GetBlog(blogId string) (*dto.Blog, error)
	CreateBlog(blog *dto.Blog) error
}


type blogInteractor struct {
	repository.IBlogRepository
}

func NewBlogInteractor(repo repository.IBlogRepository) BlogInputPort {
	return &blogInteractor{repo}
}

func (s *blogInteractor) GetBlog(blogId string) (*dto.Blog, error) {
	blogE, err := s.IBlogRepository.GetBlog(blogId)
	if err != nil {
		return nil, err
	}

	blogD := dto.NewBlog(blogE.ID, blogE.Title, blogE.Body)
	return blogD, nil
}

func (s *blogInteractor) CreateBlog(blog *dto.Blog) error {
	blogE := entity.NewInBlog(blog.Title, blog.Body)
	s.IBlogRepository.CreateBlog(blogE)
	return nil
}
