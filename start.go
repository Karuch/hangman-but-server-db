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

	fmt.Println("hello")
	
	


}