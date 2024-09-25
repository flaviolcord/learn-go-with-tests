package main

import "fmt"

func Repeat(str string) (finalStr string) {
	for i := 0; i < 4; i++ {
		finalStr += str
	}

	return
}

func main() {
	fmt.Println(Repeat("a"))
}
