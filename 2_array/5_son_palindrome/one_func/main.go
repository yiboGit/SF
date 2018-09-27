package main

func main() {

}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var (
		los, maxLen int
	)
	for i, _ := range s {
		lo, length := extendPalindrome(s, i, i)
		if length >= maxLen {
			los = lo
			maxLen = length
		}
		lo, length = extendPalindrome(s, i, i+1)
		if length >= maxLen {
			los = lo
			maxLen = length
		}
	}
	return s[los : maxLen+los]
}

func extendPalindrome(s string, j, k int) (lo, length int) {
	for j >= 0 && k < len(s) && s[j] == s[k] {
		j--
		k++
	}
	return (j + 1), (k - j - 1)
}
