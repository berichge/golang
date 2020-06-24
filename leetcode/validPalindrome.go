package main

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && ignore(s[left]) {
			left++
		}
		for left < right && ignore(s[right]) {
			right--
		}

		if (left < right) && (toLowerCase(s[left]) != toLowerCase(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func ignore(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return false
	}
	return true
}

func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}
