package main

import (
	"os"
	"fmt"
	"time"
	"math/rand"
	"io/ioutil"
	"strings"
)

// Hold the list of words that have been seen.
var wordList map[string]*Word = make(map[string]*Word)

// Holds how often this word has been used.
// It also holds a map of all following words, and how often they occur.
type Word struct {
	count int
	words map[string]int
}

// Constructor for Word.
func NewWord() *Word {
	return &Word{count: 1, words: make(map[string]int)}
}

// Take in our phrase, and parse it into words.
// This will put what it finds in the global wordList.
func ParsePhrase(s string) {
	// I assume all words are separated by spaces.
	// Normalize the strings somewhat.
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.Replace(s, "\t", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)
	var array = strings.Split(s, " ")

	// Insets the words we found into wordList.
	for i := 0; i < len(array)-1; i++ {
		_, ok := wordList[array[i]]
		// If we haven't seen this word, make a Word for it.
		if !ok {
			wordList[array[i]] = NewWord()
		}
		// Increment our current word, and the next word's count.
		wl := *wordList[array[i]]
		wl.count++
		wl.words[array[i+1]]++
	}
}

// Prints the words we have found randomly.
func PrintPhrase(w int) {
	var s = ""
	// Hack-ish way to obtain a random word to start with.
	// This works because it is undefined in what order maps will return.
	for k := range wordList {
		s = k
		break
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < w; i++ {
		// Reset some stuff.
		var q = wordList[s]
		s = ""

		for k := range q.words {
			s = k
			var a = q.words[k]
			var b = q.count
			// NOTE: this could probably be faster with a lookup table.
			if float64(b)/float64(a) * float64(rand.Int31()%100) >= 50 {
				s = k
				break
			}
		}
		// If we don't pick anything we go full random.
		if s == "" {
			for k := range q.words {
				s = k
				break
			}
		}
		fmt.Print(s + " ")
	}
}

// Starts the simulation.
func main() {
	// Read in our seed.
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	phrase := string(p)
	// Start the simulation.
	ParsePhrase(phrase)
	PrintPhrase(100)
	fmt.Println()
}
