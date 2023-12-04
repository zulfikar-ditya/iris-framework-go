package Book

import "github.com/kataras/iris/v12"

type BookStruct struct {
	Title string `json:"title"`
}

func GetListBook(ctx iris.Context) {
	books := []BookStruct{
		{Title: "Book 1"},
		{Title: "Book 2"},
		{Title: "Book 3"},
		{Title: "Book 4"},
	}

	ctx.JSON(books)
}

func CreateBook(ctx iris.Context) {
	var b BookStruct;

	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"message": "Bad Request",
			"error": err.Error(),
		})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{
		"message": "Book created",
		"data": b,
	})
}