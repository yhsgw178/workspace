package main

func main() {
	item := make([]int32, 0, 10)
	for i := 30; i <= 120; i += 15 {
		item = append(item, int32(i))
	}
	for _, v := range item {
		println("before:", v)
		println(increase(v))
	}
}

func increase(item int32) int64 {
	if item >= 90 {
		return 120
	}
	add := item + 30
	multiple := (add + 29) / 30 * 30
	return int64(multiple)
}
