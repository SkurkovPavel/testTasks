package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Problem: it is necessary to expand the array using only the "arraySwap" function

//solution->done

func main() {
	fmt.Print("Enter number: ")
	reader := bufio.NewScanner(os.Stdin)

	var basicSlice []int

	for reader.Scan() {
		breaker := reader.Text()
		i, err := strconv.Atoi(breaker)
		if err != nil {
			fmt.Print("result:")
			break
		}
		basicSlice = append(basicSlice, i)
		fmt.Print("One more number?:")
		continue
	}
	esc := false
	for !esc  {
		esc = true
		for chKey := range basicSlice {
			if chKey+1 < len(basicSlice) {
				if basicSlice[chKey] > basicSlice[chKey+1] {
					esc = false
					arraySwap(basicSlice, chKey)
					arraySwap(basicSlice, chKey+1)
					arraySwap(basicSlice, chKey)
					continue
				}
			}

		}
	}
	fmt.Println(basicSlice)
}

func arraySwap(arr []int, n int) {

	arr[n], arr[0] = arr[0], arr[n]

}


