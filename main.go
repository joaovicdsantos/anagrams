package main

import (
	"fmt"
	"github.com/joaovicdsantos/anagrams/word"
	"os"
	"time"
)

func main() {

	words := os.Args[1:]
	if len(words) == 0 || len(words) > 1 {
		fmt.Println("Usage: ./anagrams word")
		os.Exit(1)
	}

	currentWord := word.NewWord(words[0])

	fmt.Println("\n...:::: Anagrams ::::....")
	fmt.Printf("Word\t\t%s\n", currentWord.Content)
	fmt.Printf("Sliced\t\t%s\n", currentWord.Sliced)
	fmt.Printf("Repetitions\t%v\n", currentWord.Repetitions)
	fmt.Printf("Anagrams\t%d\n\n", currentWord.Anagrams)

	inicio := time.Now()

	fmt.Print("Do you want to show all possible combinations?(y/any) ")
	var res string
	fmt.Scanln(&res)
	if res == "y" {
		currentWord.ShowCombinations()
	} else {
		os.Exit(0)
	}

	fmt.Println(time.Since(inicio))
}
