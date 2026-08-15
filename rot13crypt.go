// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {

		if r[i] >= 'a' && r[i] <= 'z' {
			r [i] = 'a' + ( r[i] - 'a'  + 13 ) % 26
}	
	else if r[i] >= 'A' && r[i] <= 'Z' {
			r [i] = 'A' + ( r[i] - 'A'  + 13 ) % 26
	} 
}
return string(r)
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {

	return encrypt_rot13(s)

}