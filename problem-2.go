package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Problem-2 reverse number without strings or any build in functions

//solution->done

func main_shmain() {

	fmt.Print("Enter number: ")
	reader := bufio.NewScanner(os.Stdin)//New scanner for users number
	var number int
	for reader.Scan() {//Start scanning terminal
		breaker := reader.Text()//Get string from user
		i, err := strconv.Atoi(breaker)//Just convert string to int
		if err != nil {//If something going wrong
			fmt.Println("Oops, something went wrong")
			continue
		}
		if i == 0 {
			fmt.Println("Enter a number greater than zero:")
			continue
		}
		number = i
		break
	}
	fmt.Printf("Result: %v", inverseNumber(number))
}


func inverseNumber(n int) (i int) {
	for n != 0 {
		i *= 10
		mod := n % 10
		i += mod
		n = (n - mod) / 10
	}
	return
}