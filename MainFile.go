package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("\nWELCOME TO OUR CYPHERING TOOL!\n")
// all the variables will be assigned form the getinput function
// toEncrypt is a bool varialbe so true will encrypt and
// false will decrypt	
	
	
	toEncrypt, encoding, message := getInput() 
	// result := "" other way to create a variable
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
	} else { 
		fmt.Println("Unknow Cipher, please choose from list")
	   return 
	}	
	fmt.Println("Result:" ,result)

}

/*else if encoding == "MorseCode" {

		if toEncrypt { 
			result = encrypt_MorseCode(message)  
		} else {
			result = decrypt_MorseCode(message)
		}*/


func getInput() (toEncrypt bool, encoding string, message string){
	// i changed the var toEncrpyt to smth else because its a bool
	text := bufio.NewReader(os.Stdin)
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")
	input, _ := text.ReadString('\n')
	choice := strings.TrimSpace(input)
		
		if choice == "1" {
			fmt.Println("You chose to encrypt!\n")
			toEncrpyt = true
		} else if choice == "2" {
			fmt.Println("You chose to decrypt!\n")
			toEncrpyt = false
		} else {
			fmt.Println("Invalid choice.")
			return
		}
	
	
	fmt.Println("Select cypher (1/3):")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("3. Morse.")
	input, _ = text.ReadString('\n')
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
	return
}


var morseToLetter = map[string]rune{
	".-": 'a', "-...": 'b', "-.-.": 'c', "-..": 'd',
	".": 'e', "..-.": 'f', "--.": 'g', "....": 'h',
	"..": 'i', ".---": 'j', "-.-": 'k', ".-..": 'l',
	"--": 'm', "-.": 'n', "---": 'o', ".--.": 'p',
	"--.-": 'q', ".-.": 'r', "...": 's', "-": 't',
	"..-": 'u', "...-": 'v', ".--": 'w', "-..-": 'x',
	"-.--": 'y', "--..": 'z',
	}

var LetterToMorse = map[rune]string {
	'a': ".-",   'b': "-...", 'c': "-.-.", 'd': "-..",
	'e': ".",    'f': "..-.", 'g': "--.",  'h': "....",
	'i': "..",   'j': ".---", 'k': "-.-",  'l': ".-..",
	'm': "--",   'n': "-.",   'o': "---",  'p': ".--.",
	'q': "--.-", 'r': ".-.",  's': "...",  't': "-",
	'u': "..-",  'v': "...-", 'w': ".--",  'x': "-..-",
	'y': "-.--", 'z': "--..",
}
	

func encrypt_MorseCode(s string) string {
	s = strings.ToLower(s)
	result := ""
	for _, char := range s { 
		morse, found:= LetterToMorse[char] 
		if found { 
			result = result + morse + " "
		} else { 
			result = result + string(char)
		}
	
	}
return result
}




func decrypt_MorseCode(s string) string {

	sliced := strings.Split(s," ")
	result := ""
	for _, code := range sliced { 
		morse, found := morseToLetter[code] 
		if found {
			result = result + string(morse)
		} else {
			result = result + string(code)
		}
	}
	return result
	
}

func encrypt_rot13(s string) string {
	
	result := ""

	for _, char := range s {
    	if char >= 'a' && char <= 'z' {
        result += string('a' + (char-'a'+13)%26)
    } else if char >= 'A' && char <= 'Z' {
        result += string('A' + (char-'A'+13)%26)
    } else {
        result += string(char)
    }
}

return result

} 

// Decrypt the message with rot13
func decrypt_rot13(s string) string {

	return encrypt_rot13(s)

}


func encrypt_reverse(s string) string {
	
	result :=""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
			 
		} else if char >= 'A' && char <= 'Z'{
			result += string('Z' - (char - 'A'))
		} else {
			result += string(char)
		}
		
	}
	return result
}	



func decrypt_reverse(s string) string {
	return encrypt_reverse(s)
}


