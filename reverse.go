func encrypt_reverse(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string('z' - (char-'a')%26)
		} else if char >= 'A' && char <= 'Z' {
			result += string('Z' - (char-'A')%26)
		} else {
			result += string(char)
		}
	}

	return result
}

func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}