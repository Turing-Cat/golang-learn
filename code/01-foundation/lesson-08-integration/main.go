package main

import "fmt"

func WordCount(text []string) map[string]int {
	count := make(map[string]int)
	for _, word := range text {
		count[word]++
	}
	return count
}

func main() {
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	result := WordCount(words)
	fmt.Println(result)

	var empty []string
	result = WordCount(empty)
	fmt.Println(result)
}
