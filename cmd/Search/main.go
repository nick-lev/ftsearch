// test task: full text search - search tool

package main

import (
	"github.com/nick-lev/ftsearch/index"
	"log"
	"os"
	"path"
	"sort"
)

// structures for grade search
type kv struct {
	key   string
	value int
}

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
		log.Fatal(err)
	} else {
		log.Println("Index loaded")
	}

	if len(os.Args) < 3 {
		usage()
		log.Fatal("wrong usage")
	}

	switch os.Args[1] {
	case "-any":
		searchAny(os.Args[2:], data)
	case "-all":
		searchAll(os.Args[2:], data)
	case "-phrase":
		searchPhrase("", data)
	default:
		usage()
	}

}

func usage() {
	log.Println("Usage: [-any | -all | -phrase]")
	log.Println("\tUsage: -any   word1 word2 etc....")
	log.Println("\tUsage: -all word1 [meanean and] word2d2 etc...")
	log.Println("\tUsage: -phrase phrase to searchch from indexFileNameed files")
	return
}

func searchAny(words []string, data index.Data) {
	log.Println("Search (any words) starting:")
	for i, word := range words {
		result := []kv{}
		log.Printf("\t%v. Searching for '%v'\n", i, word)
		if _, ok := data[word]; !ok {
			log.Println("\t\tcan not find this word")
		} else {
			log.Println("\tFound in:")
			for path, count := range data[word] {
				result = append(result, kv{path, count})
			}
		}
		// will work from v 1.8
		sort.Slice(result, func(i, j int) bool {
			return result[i].value > result[j].value
		})
		for _, kv := range result {
			log.Printf("\t\t%s [%d]\n", kv.key, kv.value)
		}
	}
}

func searchAll(words []string, data index.Data) {
	log.Println("Search (all words) starting:")
	wordsMap := make(map[string]int)
	for _, w := range words {
		wordsMap[w]++
		if wordsMap[w] > 1 {
			log.Printf("duplicate search words found(will be skipped): [%v]\n", w)
		}
	}
	filesMap := make(map[string]int)
	for w, _ := range wordsMap {
		if _, ok := data[w]; ok == true {
			for path, _ := range data[w] {
				filesMap[path]++
			}
		}
	}
	count := 0
	for path, _ := range filesMap {
		if filesMap[path] == len(wordsMap) {
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
