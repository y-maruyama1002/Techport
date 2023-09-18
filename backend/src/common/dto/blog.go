package dto

type Blog struct {
	Id string `json:id`
	Title string `json:title`
	Body string `json:body`
}

func NewBlog(id string, title string, body string) *Blog {
	return &Blog {
		Id: id,
		Title: title,
		Body: body,
	}
}
