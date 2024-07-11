package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please give a one arguments")
		os.Exit(1)
	}
	arguments := os.Args[1:]
	sum := 0
	count := 0
	for _, arg := range arguments {
		m, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("An argument is not a float: ", arg)
			os.Exit(1)
		}
		sum += m
		count++
	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("Sum %d\n", sum)
		fmt.Printf("Average %.2f\n", average)
	} else {
		fmt.Println("No arguments provided")
	}
}
