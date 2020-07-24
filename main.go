package main

import (
	"math/rand"
	"time"

	"./fancynoun"
	"./wordlist"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	words := fancynoun.FancyNoun{
		Adjectives:   wordlist.WordList{WordType: "adjectives"},
		Nouns:        wordlist.WordList{WordType: "nouns"},
		Verbs:        wordlist.WordList{WordType: "verbs"},
		Prepositions: wordlist.WordList{WordType: "prepositions"},
	}
	words.LoadWords()

	for i := 0; i < 10; i++ {
		words.PrintFancyNoun()
	}
}
