package wordlist

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
)

// WordList struct contains the wordtype and word list
type WordList struct {
	WordType string
	Words    []string
}

// LoadWords loads the words from txt files
func (wl *WordList) LoadWords() {
	path := "./words/" + wl.WordType + ".txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		wl.Words = append(wl.Words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// GetRandomWord gets a random word from the word list
func (wl *WordList) GetRandomWord() string {
	number := rand.Intn(len(wl.Words))
	return wl.Words[number]
}

// NewWordList constructs a new WordList struct
func NewWordList(wordType string) *WordList {
	wl := &WordList{WordType: wordType}
	wl.LoadWords()
	return wl
}
