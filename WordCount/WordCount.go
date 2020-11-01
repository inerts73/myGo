package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount - exercise in "A Tour of Go"
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, char := range strings.Split(s," ") {
		if _, ok := m[string(char)]; ok {
			m[string(char)]++      
		} else {
			m[string(char)] = 1  
    	}   
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
