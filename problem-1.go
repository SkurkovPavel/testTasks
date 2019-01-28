package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Problem: it is necessary to sort the array using only the "arraySwap" function

//solution->done

func main() {print("Enter number: ")
	reader := bufio.NewScanner(os.Stdin)//New scanner for users numbers
	result := GetArrayFromUser(reader)
	fmt.Printf("Result: %v",result)

}

func GetArrayFromUser(reader *bufio.Scanner) []int{


	var basicSlice []int //Slice for user numbers

	for reader.Scan() { //Start scanning terminal
		breaker := reader.Text()//Get string from user
		i, err := strconv.Atoi(breaker)//Just convert string to int
		if err != nil { //If something going wrong or user finished input
			break
		}
		basicSlice = append(basicSlice, i) //Append users numbers into slice
		//fmt.Print("One more number?:") //Just say something for fun
		continue //Continue scanning
	}
	arraySort(basicSlice)

	return basicSlice //Print result
}

func arraySort(basicSlice []int)  {
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
}

func arraySwap(arr []int, n int) {
	//Swap array elements
	arr[n], arr[0] = arr[0], arr[n]
}


