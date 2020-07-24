package fancynoun

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Spacebjorn/recursivetransitionnetwork/wordlist"
)

type FancyNoun struct {
	articles         []string
	relativePronouns []string
	Adjectives       wordlist.WordList
	Nouns            wordlist.WordList
	Verbs            wordlist.WordList
	Prepositions     wordlist.WordList
}

func (fn *FancyNoun) LoadWords() {
	fn.articles = []string{"the", "an", "a"}
	fn.relativePronouns = []string{"who", "which", "that", "whose", "whom"}
	fn.Adjectives.LoadWords()
	fn.Nouns.LoadWords()
	fn.Verbs.LoadWords()
	fn.Prepositions.LoadWords()
}

func (fn *FancyNoun) getArticle() string {
	n := rand.Intn(len(fn.articles))
	return fn.articles[n]
}

func (fn *FancyNoun) getRelativePronoun() string {
	n := rand.Intn(len(fn.relativePronouns))
	return fn.relativePronouns[n]
}

// createOrnateNoun creates a random ornate noun with the structure (article) (adjective) noun
func (fn *FancyNoun) createOrnateNoun(start int) string {
	var ornateNoun string
	path := rand.Intn(3-start) + start
	switch path {
	case 0:
		ornateNoun += fn.getArticle() + " "
		ornateNoun += fn.createOrnateNoun(1)
	case 1:
		ornateNoun += fn.Adjectives.GetRandomWord() + " "
		ornateNoun += fn.createOrnateNoun(1)
	default:
		ornateNoun += fn.Nouns.GetRandomWord()
	}
	return strings.Trim(ornateNoun, " ")
}

func (fn *FancyNoun) fancyNounProNounPath() string {
	proNounPath := rand.Intn(2)
	if proNounPath == 0 {
		fancyNoun := fn.Verbs.GetRandomWord() + " "
		fancyNoun += fn.createFancyNoun()
		return fancyNoun
	}

	fancyNoun := fn.createFancyNoun() + " "
	fancyNoun += fn.Verbs.GetRandomWord() + " "
	return strings.Trim(fancyNoun, " ")
}

func (fn *FancyNoun) createFancyNoun() string {
	var fancyNoun string
	path := rand.Intn(3)
	fancyNoun += fn.createOrnateNoun(0) + " "
	if path == 0 {
		fancyNoun += fn.getRelativePronoun() + " "
		fancyNoun += fn.fancyNounProNounPath()
	} else if path == 1 {
		fancyNoun += fn.Prepositions.GetRandomWord() + " "
		fancyNoun += fn.createFancyNoun()
	}
	return strings.Trim(fancyNoun, " ")
}

func (fn *FancyNoun) PrintOrnateNoun() {
	noun := fn.createOrnateNoun(0)
	fmt.Println(noun)
}

func (fn *FancyNoun) PrintFancyNoun() {
	noun := fn.createFancyNoun()
	fmt.Println(noun)
}
