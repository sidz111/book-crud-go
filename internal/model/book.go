package model

type Book struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
	PublishedYear int    `json:"published_year,omitempty"`
}

func New(id int, name string, authorName string, publishedYear int) *Book {
	return &Book{
		ID:            id,
		Name:          name,
		AuthorName:    authorName,
		PublishedYear: publishedYear,
	}
}
