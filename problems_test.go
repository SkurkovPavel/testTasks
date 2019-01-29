package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

var testOk = `2
2
78
4
5
9
`

func TestGetArray(t *testing.T) {
	in := bufio.NewScanner(strings.NewReader(testOk))
	arr := GetArrayFromUser(in)
	out := []int{2, 2, 4, 5, 9, 78}
	result := reflect.DeepEqual(arr, out)

	if !result {
		t.Errorf("Error: Arrays are not Equal \n%v\n%v", arr, out)
	}

}
