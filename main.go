package main

import WuflyGo "github.com/bwufly/wuflyGo/framework"

func main() {
	app := WuflyGo.Classic()
	app.Map("/", func(ctx *WuflyGo.HttpContext) {
		ctx.JSON(WuflyGo.M{"info":"hello world"})
	})
	app.Run(":1688")
}


