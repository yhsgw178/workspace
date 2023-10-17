package main

import (
	"fmt"
	"sort"
)

func main() {
	data := map[string][]int32{
		"apple":   {3, 2, 1},
		"banana":  {6, 5, 4},
		"cherry":  {9, 8, 7},
		"date":    {12, 11, 10},
		"elderly": {15, 14, 13},
	}

	// スライスを逆順にソート
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	// ソートしたキーに基づいてmapを更新
	sortedData := make(map[string][]int32)
	for _, key := range keys {
		// 配列を逆順にリバース
		reversedArray := reverseIntSlice(data[key])
		sortedData[key] = reversedArray
	}

	// 結果を表示
	fmt.Println(sortedData)
}

// int32のスライスを逆順にリバースする関数
func reverseIntSlice(slice []int32) []int32 {
	reversed := make([]int32, len(slice))
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = slice[i]
	}
	return reversed
}
