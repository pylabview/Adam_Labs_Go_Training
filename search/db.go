package main

import (
	"regexp"
	"strings"
)

type Record struct {
	Title   string
	Content string
	Link    string
}

type DB struct {
	index map[string][]int // word -> doc index

}

func NewDB() *DB {
	return &DB{
		index: index,
	}
}

// Inverted index: word -> [doc id, doc id, ...]
var index map[string][]int
var re = regexp.MustCompile(`[a-z]+`)

func init() {
	/*
		for i := range dbData {
			dbData[i].Content = strings.ToLower(dbData[i].Content)
		}
	*/
	index = make(map[string][]int)
	for i, rec := range dbData {
		for _, word := range re.FindAllString(strings.ToLower(rec.Content), -1) {
			index[word] = append(index[word], i)
		}
	}
}

func (db *DB) Search(query string) []Record {
	query = strings.ToLower(query)

	indices := db.index[query]
	matches := make([]Record, len(indices))
	for i, n := range indices {
		matches[i] = dbData[n]
	}

	/*
		var matches []Record
		for _, record := range dbData {
			if strings.Contains(record.Content, query) {
				matches = append(matches, record)
			}
		}

	*/
	return matches
}
