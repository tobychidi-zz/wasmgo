package main

import (
	"syscall/js"
)

func article(target js.Value) js.Value {
	var article = js.Global().Get("document").Call("createElement", "article")
	article.Get("classList").Call("add", "column")
	article.Get("style").Set("maxWidth", "300px")
	article.Get("style").Set("minWidth", "300px")

	return article
}
