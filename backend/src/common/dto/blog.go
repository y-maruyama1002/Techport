package dto

type Blog struct {
	ID uint
	Title string `json:title`
	Body string `json:body`
}

func NewBlog(id uint, title string, body string) *Blog {
	return &Blog {
		ID: id,
		Title: title,
		Body: body,
	}
}
