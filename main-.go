package main

import "fmt"

func main(){
	//LOOPING WITH CONTUNE & BREAK
	for i := 0; i < 20; i++ {
		if i == 5{
			continue
		}
		if i > 15{
			break
		}
		fmt.Println(i)
	}
}