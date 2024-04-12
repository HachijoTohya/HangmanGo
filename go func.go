package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func selectRandomWord(difficulty string) string {
	wordFile, _ := os.Open("words.txt")
	wordScanner := bufio.NewScanner(wordFile)
	var wordsEasy []string
	var wordsMedium []string
	var wordsHard []string
	var wordList []string
	//Sort words by difficulty
	for wordScanner.Scan() {
		if len(wordScanner.Text()) <= 5 {
			wordsEasy = append(wordsEasy, wordScanner.Text())
		} else if 5 < len(wordScanner.Text()) && len(wordScanner.Text()) <= 8 {
			wordsMedium = append(wordsMedium, wordScanner.Text())
		} else {
			wordsHard = append(wordsHard, wordScanner.Text())
		}
	}
	for difficultySelect := false; !difficultySelect; {
		switch difficulty {
		case "H":
			difficultySelect = true
			if len(wordsHard) == 0 {
				fmt.Println("This difficulty does not have any valid words! Check your words.txt file.")
				difficultySelect = false
				difficulty = input()
			}
		case "M":
			difficultySelect = true
			if len(wordsMedium) == 0 {
				fmt.Println("This difficulty does not have any valid words! Check your words.txt file.")
				difficultySelect = false
				difficulty = input()
			}
		case "E":
			difficultySelect = true
			if len(wordsEasy) == 0 {
				fmt.Println("This difficulty does not have any valid words! Check your words.txt file.")
				difficultySelect = false
				difficulty = input()
			}
		default:
			fmt.Println("Invalid difficulty. Please input H, M, or E.")
			difficulty = input()
		}
	}
	wordFile.Close()
	difficulties := map[string][]string{"H": wordsHard, "M": wordsMedium, "E": wordsEasy}
	wordList = difficulties[difficulty]
	return strings.ToUpper(wordList[rand.Intn(len(wordList))])
}

func createBlank(word string) string {
	var blank string
	for _, v := range word {
		if unicode.IsLetter(v) {
			blank = blank + "_ "
		} else {
			blank = blank + string(v) + " "
		}
	}
	return blank
}

func guess(query, blank string, indexMap map[string][]int) (string, bool) {
	indexArray, ok := indexMap[query]
	switch len(query) {
	case 1:
		if ok {
			for _, v := range indexArray {
				v = v * 2
				blank = blank[:v] + query + blank[v+1:]
			}
		}
	default:
		ok = true
		fmt.Println("Please input only a single character.")
	}
	return blank, ok
}

func input() string {
	var inputString string
	for isValid := true; isValid; {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if len(scanner.Text()) != 1 {
			fmt.Println("Please only input a single alphabetic character.")
		} else if !unicode.IsLetter(rune(scanner.Text()[0])) {
			fmt.Println("Please only input a single alphabetic character.")
		} else {
			inputString = scanner.Text()
			isValid = false
		}
	}
	return strings.ToUpper(inputString)
}
func main() {
	for mainLoop := true; mainLoop; {
		//initialize game
		fmt.Println("Select a difficulty: Hard(H)/Medium(M)/Easy(E)")
		word := selectRandomWord(input())
		indexMap := map[string][]int{}
		blank := createBlank(word)
		for i, char := range word {
			indexMap[string(char)] = append(indexMap[string(char)], i)

		}
		fmt.Println("Guess the word by guessing one letter at a time. You're allowed 6 incorrect guesses.")
		fmt.Println(blank)
		//main game loop
		var lives uint8 = 6
		for strings.ReplaceAll(blank, " ", "") != word && lives != 0 {
			query := input()
			var ok bool
			blank, ok = guess(query, blank, indexMap)
			if !ok {
				lives--
			}

			fmt.Println("Lives: ", lives)
			fmt.Println(blank)
		}
		switch lives {
		case 0:
			fmt.Printf("You could not guess the word. The word was %s\n", word)
		default:
			fmt.Printf("You guessed the word! The word was %s\n", word)
		}
		fmt.Println("Would you like to play again? Y/N")

		switch strings.ToUpper(input()) {
		case ("Y"):
			mainLoop = true
		default:
			mainLoop = false
			fmt.Println("Thanks for playing!")
		}
	}
}
