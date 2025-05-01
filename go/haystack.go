package main

func newStr(haystack string, needle string) int {
	length := len(needle)

	if length == 0 {
		return 0
	}

	if length == len(haystack) {
		if haystack == needle {
			return 0
		}
	}

	for i := 0; i <= len(haystack)-length; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}

	return -1
}
