package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int
	fmt.Println("enter integers one by one, enter STOP to stop")

	for {
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToUpper(input) == "STOP" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid integer")
		} else {
			numbers = append(numbers, num)
		}
	}

	fmt.Println("Numbers:", numbers)
}
