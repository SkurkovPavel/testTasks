package main

import (
	"sort"
	"strconv"
	"strings"
)
//Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are substrings of strings of a2.
//
//#Example 1: a1 = ["arp", "live", "strong"]
//
//a2 = ["lively", "alive", "harp", "sharp", "armstrong"]
//
//returns ["arp", "live", "strong"]

func InArray(array1 []string, array2 []string) []string {
	arraymap := make(map[string]int)
	for key2 := range array2 {
		for key1, val1 := range array1 {
			if strings.Contains(array2[key2], array1[key1]) {
				arraymap[val1] = 1
			}
		}
	}
	array3 := []string{}
	for key := range arraymap {
		array3 = append(array3, key)
	}

	sort.Strings(array3)

	return array3
}

func greet() (s string) {
	s ="hello world"
	return
}

func Multiply( a , b int) (r int){
	r = a*b
	return
}

	//A bookseller has lots of books classified in 26 categories labeled A, B, ... Z.
	//Each book has a code c of 3, 4, 5 or more capitals letters.
	//The 1st letter of a code is the capital letter of the book category.
	//In the bookseller's stocklist each code c is followed by a space and by
	//a positive integer n (int n >= 0) which indicates the quantity of books of this code in stock.

func StockList(listArt []string, listCat []string) (list string) {

	listRes := make(map[string]int,0)
	if len(listCat) == 0 || len(listArt)==0{
		return
	}


	for keyArt:=range listArt {
		for keyCat := range listCat{
			if  strings.HasPrefix(listArt[keyArt],listCat[keyCat]) {
				listCounts := strings.Split(listArt[keyArt], " ")
				itemCount,_ := strconv.Atoi(listCounts[1])
				listRes[listCat[keyCat]] += itemCount

			}
		}
	}
	stringList := []string{}
	for _,val := range listCat{
		num:= strconv.Itoa(listRes[val])
		stringList = append(stringList,"("+val+" : "+num+")")
	}
	list = strings.Join(stringList, " - ")
	return
}
//Take 2 strings s1 and s2 including only letters from ato z.
// Return a new sorted string, the longest possible, containing distinct letters,

// best Practice 0.0529ms 0.0435ms
func TwoToOneBest(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
			result += char
		}
	}
	return result
}

//bad practice 0.0468ms 0.0841ms  0.0635ms
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func TwoToOne(s1 string, s2 string) (str string) {

	str = SortString(s1+s2)
	return
}

func SortString(s string) string {
	r := []rune(s)
	r = removeDuplicates(r)
	sort.Sort(sortRunes(r))
	return string(r)
}

func removeDuplicates(r []rune) (res []rune) {

	set := make(map[rune]bool)
	for _, val := range r {
		set[val] = true
	}
	for val  := range set{
		res = append(res,val)
	}
	return
}
func main(){

	println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
//You get an array of numbers, return the sum of all of the positives ones.
func PositiveSum(numbers []int) int {
	sum := 0
	for _,num := range numbers {
		if num>0 {
			sum += num
		}
	}
	return sum
}
//A hero is on his way to the castle to complete his mission.
// However, he's been told that the castle is surrounded with a couple of powerful dragons!
// each dragon takes 2 bullets to be defeated, our hero has no idea how many bullets he should carry..
// Assuming he's gonna grab a specific given number of bullets
// and move forward to fight another specific given number of dragons, will he survive?
//
//Return True if yes, False otherwise :)
// best Practice
func HeroBest (bullets, dragons int) bool {
	return bullets >= 2 * dragons
}
//bad practice
func Hero(bullets, dragons int) bool {
	dragons *= 2
	if dragons <= bullets {
		return true
	}
	return false
}