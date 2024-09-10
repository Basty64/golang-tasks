package main

import "fmt"

func main() {

	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	uniqueSet := make(map[string]bool)
	for _, str := range strings {
		uniqueSet[str] = true
	}

	fmt.Println("Множество:", uniqueSet)
}
