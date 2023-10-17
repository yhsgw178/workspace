package main

func main() {
	// 60以下は一律30とする
	item := make([]int64, 0, 10)
	for i := 30; i <= 120; i += 15 {
		item = append(item, int64(i))
	}
	for _, v := range item {
		println(thirty(v))
	}
}

func thirty(value int64) int64 {
	if value <= 60 {
		return 30
	}
	subtractedValue := value - 30
	closestMultiple := (subtractedValue / 30) * 30
	return closestMultiple
}
