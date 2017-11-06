package main

import (
	"log"
)

func lengthOfLongestSubstring(s string) (result int) {
	charLastIndex := [128]int{}
	for counter, i := 0, 0; i < len(s); i++ {
		if lastIndex := charLastIndex[int(s[i])]; lastIndex > 0 {
			for char, index := range charLastIndex { // remove last index note before index ``i``
				if index >= 1 && index <= lastIndex {
					charLastIndex[char] = -1
					counter--
				}
			}
		}

		counter++
		charLastIndex[s[i]] = i + 1 // index + 1
		if counter > result {
			result = counter
		}
	}

	return result
}

func main() {
	log.Print(lengthOfLongestSubstring("bpfbhmipx")) // 7
	log.Print(lengthOfLongestSubstring("abcabcbb"))  // 3
	log.Print(lengthOfLongestSubstring("tmmzuxt"))   // 5
}

func init() {
	log.SetFlags(log.Lshortfile)
}
