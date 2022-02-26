package main

import "fmt"

func main(){
	//LOOPING OVER STRING
	sentence := "Hello"
	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	fmt.println("   ------>>>")
	for pos, char := range sentence {
	fmt.printf("Character %c start at byte position %d/n", char, pos)
	}
}