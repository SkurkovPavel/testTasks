package main
//https://www.codewars.com/kata/first-variation-on-caesar-cipher/train/go

import "strings"

func MovingShift(s string, shift int) (moving []string) {



	//Plain alphabet [a b c d e f g h i j k l m n o p q r s t u v w x y z]
	pAlph := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ","")

	pMapAlph := make(map[string]int)
	for k,v := range pAlph{
		pMapAlph[v] = k

	}

	//Plain string
	pString := strings.Split(s,"")


	//s mkx bocod
	for x,val := range pString  {
		if pString[x] == " " {
			moving = append(moving," ")
			shift++
			continue
		}
		position := pMapAlph[val] + shift
		shift++
		for position >50{
			position =  position - 50

		}

		moving = append(moving,pAlph[position])
	}

	return
}
func DemovingShift(arr []string, shift int) (res string) {

	return
}
