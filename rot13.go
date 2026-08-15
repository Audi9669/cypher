
// Encrypt the message with rot13
func encrypt_rot13(s string) string {

	if s <= 'z' || if s >= 'a' {
	return 'a' + ( s - 'a'  + 13 ) % 26
}	
	 if s <= 'Z' || if s >= 'A' {
        return 'A' + ( s - 'A'  + 13 ) % 26
} 
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {

	return encrypt_rot13 

}
