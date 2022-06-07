package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func main() {

	http.HandleFunc("/Hello", func(res http.ResponseWriter, req *http.Request) {

		fileName := "file.txt"

		file, err := ioutil.ReadFile(fileName)

		if err != nil {

			log.Fatal(err)
		}
		// convert byteslice to string
		text := string(file)
		words := strings.Fields(text)

		// count same words
		m := make(map[string]int)
		for _, word := range words {
			m[word]++
		}

		// create and fill slice of word-count pairs for sorting by count
		wordCounts := make([]string, len(m))
		for key := range m {
			wordCounts = append(wordCounts, key)
		}

		// sort wordCount slice
		sort.Slice(wordCounts, func(i, j int) bool {
			return m[wordCounts[i]] > m[wordCounts[j]]
		})

		// get the ten most frequent words
		n := make(map[string]int)
		for index, key := range wordCounts {
			n[key] = m[key]
			fmt.Fprintf(res, "%s %d\n", key, n[key])
			if index == 9 {
				break
			}
		}
	})

	log.Fatal(http.ListenAndServe(":5100", nil))

}
