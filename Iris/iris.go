package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello go world")
		})

	iris.Get("/myjson", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"Name": "Iris",
			"Released": "13 March 2016",
			"Stars": 5525,
			})
		})
	iris.Listen(":8080")
}