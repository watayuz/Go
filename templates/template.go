/* template/index.htmlと連携する */
package main

import "github.com/kataras/iris"

type Index struct {
	Name string
}

func main() {
	iris.Config.IsDevelopment = true
	iris.Get("/", index)
	iris.Listen(":8080")
}

func index(ctx *iris.Context) {
	i := Index{}
	i.Name = "Iris"

	// デフォルトでtemplatesフォルダの直下を参照するようになっている
	ctx.MustRender("index.html", i)
}
/* irisはWebSocketもできる！！！！ */