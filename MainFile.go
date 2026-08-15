package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("WELCOME TO OUR CYPHERING TOOL!")
// all the variables will be assigned form the getinput function
// toEncrypt is a bool varialbe so true will encrypt and
//false will decrypt	
	
	
	toEncrypt, encoding, message := getInput() 

	var result string 
	
	if encoding == "rot13"{ 
		if toEncrypt {
			result = encrypt_rot13(message)
		} else {
			result = decrypt_rot13(message)
		}
	} else if encoding == "reverse" {
		if toEncrypt {
			result = encrypt_reverse(message)
		} else { 
			result = decrypt_reverse(message)
		}
	} else if encoding == "MorseCode" {
		if toEncrypt { 
			result = encrypt_MorseCode(message)  
		} else {
			result = decrypt_MorseCode(message)
		}
	} else { 
		fmt.Println("Unknow Cipher, please choose from list")
	   return 
	}	
	fmt.Println("Result:" ,result)

}


func getInput() (toEncrypt bool, encoding string, message string){
	toEncrypt := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to our amazing Cypher Tool!\n")
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	input, _ := ToEncrypt.ReadString('\n')
	choice := strings.TrimSpace(input)
		if choice == "1" {
			fmt.Println("You chose to encrypt!")
			
		} else if choice == "2" {
			fmt.Println("You chose to decrypt!")
			
		} else {
			fmt.Println("Invalid choice.")
		} 
	
	encoding := bufio.NewReader(os.Stdin)
	fmt.Println("Select cypher (1/2):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. Morse.")
	input, _ := encoding.ReadString('n')
	cypherchoice := strings.TrimSpace(input)
		if cypherchoice == "1" {
			fmt.Println("You chose ROT13")
			
		} else if cypherchoice == "2" {
			fmt.Println("You chose Reverse")
			
		} else if cypherchoice == "3" {
			fmt.Println("You chose Morse")
			
		} else {
			fmt.Println("Invalid choice.")
		}
}



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


func decrypt_MorseCode(s string) string {
	morseToLetter := map[string]rune{
		".-": 'a', "-...": 'b', "-.-.": 'c', "-..": 'd',
		".": 'e', "..-.": 'f', "--.": 'g', "....": 'h',
		"..": 'i', ".---": 'j', "-.-": 'k', ".-..": 'l',
		"--": 'm', "-.": 'n', "---": 'o', ".--.": 'p',
		"--.-": 'q', ".-.": 'r', "...": 's', "-": 't',
		"..-": 'u', "...-": 'v', ".--": 'w', "-..-": 'x',
		"-.--": 'y', "--..": 'z',
	}

	result := ""
}


// Encrypt the message with rot13
func encrypt_rot13(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i] <= 'z' {
			r[i] = 'a' + ( r[i] - 'a'  + 13 ) % 26
		}	else if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = 'A' + ( r[i] - 'A'  + 13 ) % 26
		} 
	}

	return string(r)
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {

	return encrypt_rot13(s)

}


func encrypt_reverse(s string) string {
	
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return 'z' - (char - 'a') % 26
		} else if char >= 'A' && char <= 'Z'{
			return 'Z' - (char - 'A') % 26
		} else {
			return char
		}
		result := char + strconv.Itoa(char)
	}
}



func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}


