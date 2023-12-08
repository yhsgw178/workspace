package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]string{"a3": "a", "a5": "e", "a2": "b", "a1": "a", "a4": "d"}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
}
