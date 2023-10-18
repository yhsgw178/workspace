package flag

import "fmt"

func flag() {
	isFlag1 := false
	num := 3
	for !isFlag1 {
		if isFlag2(num) {
			isFlag1 = true
			break
		}
		fmt.Println("after break")
	}
	fmt.Println("after for")
}

func isFlag2(num int) bool {
	if num == 2 {
		return true
	}
	return false
}
