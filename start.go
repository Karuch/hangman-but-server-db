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
	var guessed_true bool = false
	guess_loop:
	for 0 < 1 {
		scanner.Scan()
		guess := scanner.Text()
		fmt.Println(guess, "<-")
		if len(guess) > 1 {
			if !strings.EqualFold(guess, word) {
				//fmt.Printf("`%v` is incorrect word, try again\n", guess) ***************
				continue 
			} else {
				guessed_true = true
				break guess_loop
			}
		}

		guess = strings.ToUpper(guess)
		var possible_letter bool = false
		for index, _ := range possibilities_list {
			if guess == possibilities_list[index] {
				fmt.Println("yes")
				possible_letter = true
				break guess_loop
			} else {
				continue
			}										
		}

		if !possible_letter {
			fmt.Println("`%v` is not a valid character. (english letters only)", guess)
		}
	}
	return guess, guessed_true
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
	var guessed_true bool
	guess, guessed_true = input_checker(word, scanner)
	if guessed_true {
		fmt.Println("TRUE WORD!!!")
	}

	//initiallize the secret word into list and check the guess against it
	letters_list := strings.Split(word, "")
	letters_list_blank := make([]string, len(letters_list))
	copy(letters_list_blank, letters_list)
	
	for index, _ := range letters_list_blank {
		letters_list_blank[index] = "_"
	}

	for index, _ := range letters_list {
		if strings.EqualFold(letters_list[index], guess) {
			letters_list_blank[index] = letters_list[index]
		}
	}

	if guessed_true {
		fmt.Printf("Success, The word was `%v`!\n", strings.Join(letters_list, ""))
	}
	fmt.Println(letters_list)
	fmt.Println(letters_list_blank)
}