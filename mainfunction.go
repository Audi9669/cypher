package main 
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
	fmt.Println("Result:" result)

}