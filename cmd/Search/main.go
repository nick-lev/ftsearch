// test task: full text search - search tool

package main

import (
	"github.com/nick-lev/ftsearch/index"
	"log"
	"os"
	"path"
	"sort"
)

var dataFile string = "index.db"

func main() {

	if ex, err := os.Executable(); err != nil {
		log.Fatal(err)
	} else {
		dir := path.Dir(ex)
		dataFile = dir + "/" + dataFile
	}

	data := make(index.Data)

	if err := index.Load(dataFile, data); err != nil {
		if os.IsNotExist(err) {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Index loaded")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: Search [ word1 word2 ... | -all word1 word2 ...| -phrase \"phrase to search\"]")
	}

	switch os.Args[1] {
	case "-all":
		searchAll(os.Args[2:], data)
	case "-phrase":
		searchPhrase("", data)
	default:
		searchAny(os.Args[1:], data)
	}
}

func searchAny(words []string, data index.Data) {
	log.Println("Search (any words) starting:")
	type pc struct {
		Path  string
		Count int
	}
	words = removeDup(words)
	for i, word := range words {
		var result []pc
		log.Printf("\t%v. Searching for '%v'\n", i, word)
		if _, ok := data[word]; !ok {
			log.Println("\t\tcan not find this word")
		} else {
			log.Println("\tFound in:")
			for path, count := range data[word] {
				result = append(result, pc{Path: path, Count: count})
			}
		}
		// will work from v 1.8
		sort.Slice(result, func(i, j int) bool {
			return result[i].Count > result[j].Count
		})
		for _, r := range result {
			log.Printf("\t\t%s [%d]\n", r.Path, r.Count)
		}
	}
}

func searchAll(words []string, data index.Data) {
	log.Println("Search (all words) starting:")
	words = removeDup(words)
	f2p := make(map[string]int)
	for _, w := range words {
		if _, ok := data[w]; ok {
			for path, _ := range data[w] {
				f2p[path]++
			}
		}
	}
	count := 0
	for path, _ := range f2p {
		if f2p[path] == len(words) {
			log.Printf("\t\twords found at: %v\n", path)
			count++
		}
	}
	if count == 0 {
		log.Println("No match found!")
	}
}

func searchPhrase(phrase string, data index.Data) {
	log.Println("Search (phrase) starting:")
}

func removeDup(words []string) []string {
	wordsUniq := make(map[string]int)
	for _, w := range words {
		wordsUniq[w]++
	}
	var res []string
	for w, _ := range wordsUniq {
		res = append(res, w)
	}
	return res
}
