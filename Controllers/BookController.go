package controllers

type BookController struct {
	//
}

type Book struct {
	Title string `json:"title"`
}

func (c *BookController) Get() []Book {
	books := []Book{
		{Title: "Book 1"},
		{Title: "Book 2"},
		{Title: "Book 3"},
		{Title: "Book 4"},
	}

	return books
}

func (c *BookController) Post(b Book) Book {
	return Book{
		Title: b.Title,
	}
}