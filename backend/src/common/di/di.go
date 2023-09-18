package di

import (
	"github.com/y-maruyama1002/Techport/controller"
	"github.com/y-maruyama1002/Techport/interface/database"
	"github.com/y-maruyama1002/Techport/usecase"
	"gorm.io/gorm"
)

func InitBlog(db *gorm.DB) controller.IBlogController {
	r := database.NewBlogRepository(db)
	s := usecase.NewBlogInteractor(r)
	return controller.NewBlogController(s)
}
