package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id           string    `json:"id"`
	TitleText    TitleText `json:"titleText"`
	TitleType    TitleType `json:"titleType"`
	PrimaryImage Image     `json:"primaryImage"`
}

type TitleText struct {
	Text string `json:"text"`
}

type TitleType struct {
	Text string `json:"text"`
}
type Image struct {
	Url string `json:"url"`
}

func fetchMovies() Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://moviesdatabase.p.rapidapi.com/titles/search/keyword/love", nil)

	if err != nil {
		// log.Fatalln(err)
		js.Global().Get("console").Call("log", err)
	}

	header := req.Header

	header.Set("X-RapidAPI-Host", "moviesdatabase.p.rapidapi.com")
	header.Set("X-RapidAPI-Key", "38a48feaa9mshaf0b6e964d158fcp14da0ajsn901265d4d1b0")

	query := req.URL.Query()

	query.Set("info", "mini_info")
	query.Set("limit", "50")
	query.Set("page", "2")
	query.Set("titleType", "movie")

	req.URL.RawQuery = query.Encode()

	// log.Printf("query", query)

	res, err := client.Do(req)

	if err != nil {
		// log.Fatalln(err)
		js.Global().Get("console").Call("log", err)
	}

	//Create a variable of the same type as our model

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// log.Fatal(err)
		js.Global().Get("console").Call("log", err)
	}

	var responseObject Response
	//Decode the data
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
