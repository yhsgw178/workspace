package main

import (
	"fmt"
	"sort"
)

func main() {
	data := map[string][]int32{
		"a": {1, 2, 3},
		"b": {4, 5, 6},
		"c": {7, 8, 9},
		"d": {10, 11, 12},
		"e": {13, 14, 15},
	}

	// キーを逆順にソート
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	// キーを逆順にリバースし、各配列も逆順にリバース
	reversedData := make(map[string][]int32)
	for _, key := range keys {
		reversedArray := reverseIntSlice(data[key])
		reversedData[key] = reversedArray
	}

	// 結果を表示
	for i := range keys {
		fmt.Println(keys[i], reversedData[keys[i]])
	}
}

func reverseIntSlice(slice []int32) []int32 {
	reversed := make([]int32, len(slice))
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = slice[i]
	}
	return reversed
}
