package main

import (
	"fmt"
	"strings"
)

func main() {
	var choice string
	var iter int = 0
	for iter < 1 {
		fmt.Println("--Welcome to hangman (CLI Version)--\nto start type `start`\nto quit type `quit`")
		fmt.Scan(&choice)
		if strings.EqualFold(choice, "start") {
			break
		} else if strings.EqualFold(choice, "quit") {
			fmt.Println("exit...")
			return 
		} else {
			fmt.Printf("`%v` is wrong option, only `start` or `quit` is available", choice)
		}
	}

	//if start chosen will execute the following

	var word string = "test"
	letters_list := strings.Split(word, "")
	letters_list_blank := make([]string, len(letters_list))
	copy(letters_list_blank, letters_list)
	
	for index, _ := range letters_list_blank {
		letters_list_blank[index] = "_"
	}

	fmt.Println(letters_list)
	fmt.Println(letters_list_blank)

	var guess string
	var possibilities_list = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	fmt.Scan(&guess)
	if strings.EqualFold(guess, word) {
		fmt.Println("Correct!")
	}
	for index, _ := range possibilities_list {
		if  len(guess) > 1 {
			fmt.Printf("`%v` is not a valid character (contains more then 1).", guess)
			break
		} else if strings.EqualFold(guess, possibilities_list[index]) {
			fmt.Println("yes")
			break
		} else {
			fmt.Printf("`%v` is not a valid character. (english letters only)", guess)
			break
		}
	}
	


}