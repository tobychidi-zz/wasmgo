package main

import (
	"strings"
)

func filterMovies(query string, dataList []Result) []Result {
	var filteredList []Result

	for i := 0; i < len(dataList); i++ {
		item := dataList[i]
		title := strings.ToLower(item.TitleText.Text)
		if strings.Contains(title, query) {
			filteredList = append(filteredList, item)
		}
	}
	return filteredList
}
