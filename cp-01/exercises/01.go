package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("please provide at least one argument")
		os.Exit(1)
	}
	arguments := os.Args[1:]
	//test := os.Args[1]
	fmt.Println("isi test", arguments)
	var sum int
	for _, arg := range arguments {
		m, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("An argument is not an integer: ", arg)
			os.Exit(1)
		}
		sum += m
	}
	fmt.Println("Sum:", sum)
}
