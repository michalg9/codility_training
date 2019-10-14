package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(561892))
}

func Solution(N int) int {

	s := strconv.FormatInt(int64(N), 2)
	maxL := 0
	currL := 0
	p := false
	for i := 0; i < len(s); i++ {
		if p == false && s[i] == '0' {
			p = true
			currL++
			continue
		}

		if p == true {
			if s[i] == '0' {
				currL++
			} else {
				if currL > maxL {
					maxL = currL
				}
				currL = 0
				p = false
			}

			continue
		}
	}

	return maxL
}
