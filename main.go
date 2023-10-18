package main

import "fmt"

type Foo struct {
	ID    string
	index int32
}

type Parent struct {
	foo []Foo
}

func main() {
	parent := Parent{
		foo: []Foo{
			{
				ID:    "a",
				index: 1,
			},
			{
				ID: "b",
			},
		},
	}
	ids := []string{"a", "b"}
	fooMap := make(map[string][]int32)
	for i := range ids {
		for j := range parent.foo {
			if ids[i] == parent.foo[j].ID {
				fooMap[ids[i]] = append(fooMap[ids[i]], parent.foo[j].index)
			}
		}
	}
	fmt.Print(fooMap)
}
