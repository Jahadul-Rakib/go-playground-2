package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("please input your age: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Println("Failed to parse integer",err.Error())
	} else {
		fmt.Println("you typed: ", input)
	}

}
