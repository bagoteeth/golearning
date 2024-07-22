package main

import (
	"fmt"
)

func longestPalindrome(s string) ([][]bool, string) {
	arr := []rune(s)
	maxlen := 0
	res := s[0:1]
	dp := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]bool, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j+i < len(arr); j++ {
			if i == 0 {
				dp[j][j+i] = true
			} else if i == 1 {
				dp[j][j+i] = arr[j] == arr[j+i]
				if dp[j][j+i] && 1 > maxlen {
					maxlen = 1
					res = s[j : j+i]

				}
			} else {
				dp[j][j+i] = dp[j+1][j+i-1] || arr[j] == arr[j+i]
				if dp[j][j+i] && i > maxlen {
					maxlen = i
					res = s[j : j+i]
				}
			}
		}
	}
	return dp, res
}

func main() {
	_, b := longestPalindrome("asdfgfddfsaovsadfusadlfsadga")
	fmt.Println(b)
}
