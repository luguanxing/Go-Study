package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLen := 0
	for i, ch := range []byte(s) {
		lastI, exist := lastOccurred[ch]
		// 该字符ch出现过而且在start后（遇到重复字符）时，需要把start移到该ch后一位使得ch不重复（消除重复）
		if exist && lastI >= start {
			start = lastI + 1
		}
		// 记录最大长度
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		// 记录该字符位置
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcsbx"))
}
