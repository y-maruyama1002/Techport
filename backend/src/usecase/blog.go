package usecase

import (
	"github.com/y-maruyama1002/Techport/common/dto"
	"github.com/y-maruyama1002/Techport/domain/repository"
)

type BlogInputPort interface {
	GetBlog(blogId string) (*dto.Blog, error)
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

	blogD := dto.NewBlog(blogE.Id, blogE.Title, blogE.Body)
	return blogD, nil
}
