# ftsearch

ftsearch is just test task for text file full text search/indexing

## Installation

[Go][] version 1.8 or higher is required. Install or update ftsearch using
the 'go get' command:
    
    go get -u github.com/nick-lev/ftsearch/cmd/Index
    go get -u github.com/nick-lev/ftsearch/cmd/Search

## Usage

    Index path/2/file
    Search [ word1 word2 ... | -all word1 word2 ...| -phrase "phrase to search"]

## TODO
    
    Search -phrase functionality 
