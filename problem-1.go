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
	reader := bufio.NewScanner(os.Stdin)//New scanner for users numbers

	var basicSlice []int //Slice for user numbers

	for reader.Scan() { //Start scanning terminal
		breaker := reader.Text()//Get string from user
		i, err := strconv.Atoi(breaker)//Just convert string to int
		if err != nil { //If something going wrong or user finished input
			fmt.Print("result:")
			break
		}
		basicSlice = append(basicSlice, i) //Append users numbers into slice
		fmt.Print("One more number?:") //Just say something for fun
		continue //Continue scanning
	}
	esc := false //Flag to exit an infinite loop
	for !esc { //Star infinite loop
		esc = true //Set the flag to exit the loop when we finish sorting
		for chKey := range basicSlice { //Start sorting
			if chKey+1 < len(basicSlice) { // we need to check "out of range" error
				if basicSlice[chKey] > basicSlice[chKey+1] {//We need to check n element bigger then n+1
					esc = false //Set the flag which shown that we had changed something in array
					arraySwap(basicSlice, chKey)  //sort
					arraySwap(basicSlice, chKey+1) //sort
					arraySwap(basicSlice, chKey) //sort
					continue
				}
			}

		}
	}
	fmt.Println(basicSlice) //Print result
}

func arraySwap(arr []int, n int) {
	//Ыwap array elements
	arr[n], arr[0] = arr[0], arr[n]

}


