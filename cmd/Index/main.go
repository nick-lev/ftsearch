// test task: full text search - index tool
package main

import (
	"fmt"
	"github.com/nick-lev/ftsearch/index"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var dataFile string = "index.db"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: Index path/2/file")
	}

	if ex, err := os.Executable(); err != nil {
		log.Fatal(err)
	} else {
		dir := path.Dir(ex)
		dataFile = dir + "/" + dataFile
	}

	data := make(index.Data)

	if err := index.Load(dataFile, data); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
		log.Print(err)
	} else {
		log.Print("Index loaded")
	}

	if err := parseFile(os.Args[1], data); err != nil {
		log.Fatal(err)
	}
	log.Print("File parsed")

	if err := index.Save(dataFile, data); err != nil {
		log.Fatal(err)
	}
	log.Print("Index saved")
}

func parseFile(path string, data index.Data) error {

	source, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	mime := http.DetectContentType(source)
	if !strings.HasPrefix(mime, "text/") {
		return fmt.Errorf("file is not text")
	}

	str := string(source)
	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(str, -1)

	// init internal map as main map value
	for word, _ := range data {
		if _, ok := data[word][path]; ok {
			delete(data[word], path)
		}
	}
	for _, word := range words {
		if _, ok := data[word]; !ok {
			data[word] = make(map[string]int)
		}
		data[word][path]++
	}
	return nil
}
