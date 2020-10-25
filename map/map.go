package main

import "fmt"

func main() {
	var a []int = []int{1,3,4,56,7,12,4,6,4}
	var mp map[int]int = map[int]int {}

	for _, elem := range a {
		_, ok := mp[elem]
		if ok {
			mp[elem]++
		} else {
			mp[elem] = 1
		}
	}

	for key, elem := range mp {
		if elem > 1 {
			fmt.Println(key)
		}
	}
}