package main

import "fmt"

func main() {
	test := []string{"Hello World", "   fly me   to   the moon  ", "luffy is still joyboy", "Test"}
	for _, s := range test {
		index := lengthOfLastWord(s)
		fmt.Println(index)
	}
}
