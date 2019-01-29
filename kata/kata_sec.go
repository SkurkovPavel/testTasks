package main

import "strings"


func MovingShift(s string, shift int) (sArr []string) {

	sArr = strings.Split(s,"")
	alphabet := getAlphabet()
	lensArr := len(sArr)
	for key := range sArr  {
		position := key + shift
		if sArr[key] == " " {
			shift++
			continue

		}
		if position < lensArr{

			sArr[key] = alphabet[position]

		} else {
			sArr[key] = alphabet[position-lensArr]
		}
		shift++
	}

	return
}
func DemovingShift(arr []string, shift int) (res string) {

	return
}
func getAlphabet() []string {

	return strings.Split("abcdefghijklmnopqrstuvwxyz","")
}