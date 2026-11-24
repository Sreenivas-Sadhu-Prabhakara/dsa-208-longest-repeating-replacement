package main

import "fmt"

// CharacterReplacement returns the length of the longest substring that can
// be made uniform by changing at most k characters. Sliding window: a
// window is viable while (size − count of its most frequent letter) ≤ k.
func CharacterReplacement(s string, k int) int {
	var count [26]int
	maxFreq := 0 // historical max window frequency; never decreased (see README)
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}
		if right-left+1-maxFreq > k {
			count[s[left]-'A']--
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func main() {
	fmt.Println(CharacterReplacement("ABAB", 2))
	fmt.Println(CharacterReplacement("AABABBA", 1))
}
