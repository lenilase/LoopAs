package main

import "fmt"

const phi = 3.14

func main() {
	var r, t, luas float64
	fmt.Println("Input Jari-jari : ")
	fmt.Scanf("%v", &r)
	fmt.Println("Input tinggi : ")
	fmt.Scanf("%v", &t)

	const pi float64 = 3.14
	fmt.Println(pi)

	luas = 2 * 3.14 * r * (r + t)
	fmt.Println("Luas Permukaan Tabung : ", luas)
}
