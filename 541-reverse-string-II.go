package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverseStr(s string, k int) string {
	buf := make([]byte, 0, len(s))
	for offset := 0; offset < len(s); offset += 2 * k {
		half := min(offset+k, len(s))
		for p := half - 1; p >= offset; p-- {
			buf = append(buf, s[p])
		}

		tail := min(offset+2*k, len(s))
		for i := half; i < tail; i++ {
			buf = append(buf, s[i])
		}
	}

	return string(buf)
}
