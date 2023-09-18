package entity

func NewInBlog(title string, body string) *Blog {
	return &Blog {
		Title: title,
		Body: body,
	}
}
