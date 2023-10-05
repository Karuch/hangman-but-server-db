package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"reflect"
)


func input_checker(word string, scanner *bufio.Scanner) (string, bool) {
	var guess string
	var possibilities_list = [27]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var guessed_whole_word bool = false
	var possible_letter bool = false
	guess_loop:
	for 0 < 1 {
		scanner.Scan()
		guess := scanner.Text()
		if len(guess) > 1 {
			if !strings.EqualFold(guess, word) {
				//fmt.Printf("`%v` is incorrect word, try again\n", guess) ***************
				continue guess_loop
			} else {
				guessed_whole_word = true
				break guess_loop
			}
		}

		guess = strings.ToUpper(guess)
		for index, _ := range possibilities_list {
			if guess == possibilities_list[index] {
				fmt.Println("yes")
				possible_letter = true
				return guess, guessed_whole_word
			} else {
				continue
			}										
		}

		if !possible_letter {
			fmt.Println("`%v` is not a valid character. (english letters only)", guess)
		}
	}
	return guess, guessed_whole_word
}


func main() {
	scanner := bufio.NewScanner(os.Stdin) // later will be pass by reference to the input checker
	for 0 < 1 {
		fmt.Println("--Welcome to hangman (CLI Version)--\nto start type `start`\nto quit type `quit`")
		scanner.Scan()
		fmt.Println("Type of x:", reflect.TypeOf(scanner))
		choice := scanner.Text()
		if strings.EqualFold(choice, "start") {
			break
		} else if strings.EqualFold(choice, "quit") {
			fmt.Println("exit...")
			return 
		} else {
			fmt.Printf("`%v` is wrong option, only `start` or `quit` is available\n", choice)
		}
	}

	//if start chosen will execute the following
	var word string = "Test next"

	//checking if guess input is allowed
	var guess string
	var guessed_whole_word bool
	//initiallize the secret word into list
	letters_list := strings.Split(word, "")
	letters_list_blank := make([]string, len(letters_list))
	copy(letters_list_blank, letters_list)
	for index, _ := range letters_list {
		if letters_list[index] == " " { //case when space used in secret word
			letters_list_blank[index] = " "
		} else {
			letters_list_blank[index] = "_"
		}
	}

	//trys loop
	for 0 < 1 {
		guess, guessed_whole_word = input_checker(word, scanner)
		
		for index, _ := range letters_list {
			if strings.EqualFold(letters_list[index], guess) {
				letters_list_blank[index] = letters_list[index]
			}
		}

		fmt.Println(letters_list)
		fmt.Println(letters_list_blank)
		if !guessed_whole_word {
			for index, value := range letters_list_blank { //not work as it suppose
				if value == "_" {
					fmt.Println("rrrr") 
					break 
				} else if value == " " {
					continue
				}
				if index == len(letters_list_blank) - 1 { //wait for it to check the whole list for _
					guessed_whole_word = true
				}
			}
		}
		if guessed_whole_word {
			fmt.Printf("Success, The word was `%v`!\n", strings.Join(letters_list, ""))
			break
		} 
	}
}