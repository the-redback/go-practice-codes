package main

import (
	"fmt"
)

func assertEq(test, ans [][]string) bool {
	if len(test) != len(ans) {
		return false
	}
	for i:=0; i<len(test);i++ {
		if len(test[i]) != len(ans[i]) {
			return false
		}
		for j:=0; j<len(test[i]); j++ {
			if test[i][j] != ans[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	str1 := [][]string{{"1", "2", "3"}, {"1", "2", "3"}, {"1", "2", "3"}}
	str2 := [][]string{{"1", "2", "3"}, {"1", "2", "3"}, {"1", "2", "3"}}
	str3 := [][]string{{}}

	fmt.Println(assertEq(str1, str2)) // answer is true
	fmt.Println(assertEq(str1, str3)) //answer is false
}
