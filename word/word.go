package word

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Word - a word class
type Word struct {
	Content     string
	Sliced      []string
	Repetitions map[string]int
	Anagrams    int
}

// NewWord - create a new word
func NewWord(content string) *Word {
	slicedWord := split(content)
	rep := repetitions(slicedWord)
	angr := anagrams(content, rep)
	return &Word{
		content,
		slicedWord,
		rep,
		angr,
	}
}

func split(content string) []string {
	return strings.Split(content, "")
}

func repetitions(slicedWord []string) map[string]int {
	// Copy
	slicedWord2 := make([]string, len(slicedWord))
	copy(slicedWord2, slicedWord)

	repetitions := make(map[string]int)

	for key, letter := range slicedWord {
		repetitions[letter] = 1
		for key2, letter2 := range slicedWord2 {
			if key != key2 && letter == letter2 {
				slicedWord = append(slicedWord[:key2], slicedWord[key2:]...)
				repetitions[letter] += 1
			}
		}
	}

	return repetitions
}

func anagrams(content string, repetitions map[string]int) int {
	result := factorial(len(content))
	for _, r := range repetitions {
		result /= factorial(r)
	}
	return result
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// ShowCombinations - Calculates and shows all anagrams
func (word *Word) ShowCombinations() {
	var wg sync.WaitGroup

	calc := func(content string, firstLetter string, wg *sync.WaitGroup) {
		defer wg.Done()
		var numRandom int
		possibilities := anagrams(content, repetitions(split(content)))
		contentSplit := split(content)
		anagram := ""
		count := 0
		exist := false
		countedWords := make([]string, possibilities)
		for count < possibilities {
			numRandom = rand.Intn(len(contentSplit))
			anagram = fmt.Sprintf("%s%s", anagram, contentSplit[numRandom])
			contentSplit = append(contentSplit[:numRandom], contentSplit[numRandom+1:]...)
			if len(anagram) == len(content) {
				anagram = fmt.Sprintf("%s%s", firstLetter, anagram)
				for _, word := range countedWords {
					if word == anagram {
						exist = true
					}
				}
				if !exist {
					fmt.Println(anagram)
					count += 1
					countedWords = append(countedWords, anagram)
				}
				exist = false
				anagram = ""
				contentSplit = split(content)
			}
		}
	}

	for letter := range word.Repetitions {
		content, err := findAndRemoveOneLetter(word.Content, letter)
		if err != nil {
			os.Exit(1)
		}
		wg.Add(1)
		go calc(content, letter, &wg)
	}

	wg.Wait()

}

func findAndRemoveOneLetter(content string, letter string) (string, error) {
	sliceWord := split(content)
	for key, value := range sliceWord {
		if value == letter {
			sliceWord = append(sliceWord[:key], sliceWord[key+1:]...)
			return strings.Join(sliceWord, ""), nil
		}
	}
	err := errors.New("not found")
	return "", err
}
