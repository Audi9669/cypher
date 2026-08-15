

// this func encrypts words with morsecode
func encrypt_MorseCode(s string) {
	s =strings.ToLower(s)
	morseCode := map[rune]string{
	'a': ".-",   'b': "-...", 'c': "-.-.", 'd': "-..",
	'e': ".",    'f': "..-.", 'g': "--.",  'h': "....",
	'i': "..",   'j': ".---", 'k': "-.-",  'l': ".-..",
	'm': "--",   'n': "-.",   'o': "---",  'p': ".--.",
	'q': "--.-", 'r': ".-.",  's': "...",  't': "-",
	'u': "..-",  'v': "...-", 'w': ".--",  'x': "-..-",
	'y': "-.--", 'z': "--..",
}
result := ""

 }