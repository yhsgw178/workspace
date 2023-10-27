package main

import "fmt"

func main() {
	ids := []string{"a", "b", "c", "d", "e"}
	data := ids[:len(ids)-1]
	fmt.Println(data)
}
