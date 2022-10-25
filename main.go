package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

func server(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	str := "Demo Demo Test test, demo string for this code to check the frequency of top ten words.."

	count := make(map[string]int)

	for _, word := range strings.Fields(str) {
		count[word]++
	}

	words := make([]string, 0, len(count))
	for i := range count {
		words = append(words, i)
	}

	sort.Slice(words, func(i, j int) bool {
		return count[words[i]] > count[words[j]]
	})

	for i := 0; i < 10 && i != len(words); i++ {
		fmt.Fprintf(res, "Word  %s            count:%d\n", words[i], count[words[i]])

	}
	io.WriteString(res, "Success")

}
func main() {
	http.HandleFunc("/", server)
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal(err)
	}
}
