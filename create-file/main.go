package main

import (
	"os"
	"strconv"
)

func main() {
	err := os.WriteFile("go.txt", []byte(strconv.Itoa(number)), 0600)
	if err != nil {
		panic(err)
	}

	numbers := []int{1, 2, 3, 4, 5}

	f, _ := os.Create("go_create.txt")

	for _, number := range numbers {
		_, _ = f.WriteString(strconv.Itoa(number) + "\n")
	}
}
