package main

import (
	"syscall/js"

	"github.com/aymerick/raymond"
)

var data = fetchMovies()

var htmlSource = `
	<div class="card">
   <div class="card-image">
	{{#if primaryImage.url}}
		<img src="{{ primaryImage.url }}" class="img-responsive" style="width: 300px; height: 400px; object-fit: cover;">
	{{else}}
		<img src="https://via.placeholder.com/300x400/?text=Movie"  class="img-responsive" style="width: 300px; height: 400px; object-fit: cover;">
	{{/if}}
      
   </div>
   <div class="card-header">
      <h5 class="card-title">{{titleText.text}}</h5>
      <p class="card-subtitle text-gray">{{titleType.text}}</p>
   </div>
</div>
	`

var tpl, err = raymond.Parse(htmlSource)

func renderData(data []Result, tpl *raymond.Template) {
	var target = js.Global().Get("document").Call("querySelector", "#target")
	target.Set("innerHTML", "")

	for _, item := range data {
		atcl := article(target)

		result, err := tpl.Exec(item)
		if err != nil {
			panic(err)
		}
		atcl.Set("innerHTML", result)
		target.Call("appendChild", atcl)
	}
}

func searchNow(this js.Value, i []js.Value) any {
	query := i[0].String()

	results := filterMovies(query, data.Results)
	renderData(results, tpl)

	// js.Global().Get("console").Call("log", "test from wasm")

	return nil
}

func main() {
	c := make(chan bool)

	if err != nil {
		panic(err)
	}

	js.Global().Set("searchNow", js.FuncOf(searchNow))

	renderData(data.Results, tpl)

	<-c
}
