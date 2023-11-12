package main

import (
	"fmt"

	"github.com/Ade1aida/contains/pkg/contains"
)

func main() {

	ourFil := "C:/Users/adeli/OneDrive/Рабочий стол/учёба/go/test.txt"
	text := "cat"
	result, err := contains.Contain(ourFil, text)

	if err != nil {
		fmt.Print(err)
	} else {
		print(result)
	}
}
