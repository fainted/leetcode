package main

// log.Printf("[%s]", reverseString(""))           // ""
// log.Printf("[%s]", reverseString("1"))          // "1"
// log.Printf("[%s]", reverseString("1234567890")) // "0987654321"
func reverseString(s string) string {
	buf := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		buf = append(buf, s[i])
	}

	return string(buf)
}
