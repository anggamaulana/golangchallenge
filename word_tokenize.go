package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type Results struct {
	Data []WordCount `json:"data"`
}

func main() {
	http.HandleFunc("/", wordTokenizerAndCount)
	fmt.Println("Server start...")
	http.ListenAndServe(":8080", nil)
}

func wordTokenizerAndCount(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(b)

	fields := strings.FieldsFunc(text, func(r rune) bool {

		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
	})

	wordsCount := make(map[string]int)

	for _, field := range fields {

		wordsCount[field]++
	}

	keys := make([]string, 0, len(wordsCount))

	for key := range wordsCount {

		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {

		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})

	var results []WordCount

	for idx, key := range keys {

		results = append(results, WordCount{
			Word:  key,
			Count: wordsCount[key],
		})

		if idx == 9 {
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	jsonEncode, err := json.Marshal(Results{Data: results})

	if err == nil {
		w.Write(jsonEncode)
	} else {
		log.Fatalln(err)
	}

}
