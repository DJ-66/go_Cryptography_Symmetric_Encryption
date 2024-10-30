package main

import (
	"crypto/aes"
	//Using AES alone, you can only encrypt/decrypt
	// data that is 16bytes long or longer(which is the size of an AES block)
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// Using AES + GCM (combined mode) aleviates the AES size limitation
// say if we wanted to use random length user generated strings (often less than 16bytes)
// Aes + GCM also gives us message authentication (integrity) = Authenticated Encyption

var (
	// We're using a 32 byte long seed_String
	// seed_String will be used to programmically generate a random 32byte long salt string
	seed_String string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
	salt = seed_String
)
// The first 5 lines in main()
// easily accomplish this task
/*
	seed_String := make([]byte, 32)
	_, err := rand.Read(seed_String)
	if err != nil {
		fmt.Print(err)
	}
	*/

	// The rest of the program uses our randomly generated 32 byte salt string to encypt/decrypt plaintext
	func encrypt(plaintext string) string {
		aes, err := aes.NewCipher([]byte(salt))
		if err !=nil {
			panic(err)
		}
	// using gcm along with aes allows us to encrypt/decrypt random length strings 
	// which can be shorter than 16 bytes long (which the size of an AES block)
	gcm, err := cipher.NewGCM(aes)
	if err !=nil {
		panic(err)
	}	
	// we need a 12 byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// a nonce should always be randomly generated for every encyption.
	nonce := make ([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err !=nil {
		panic(err)
	}
	
	// ciphertext here is actually nonce + ciphertext
	// So that when we decypt, just knowing the nonce size
	// will be enough to seperate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext)
	}

	func decrypt(ciphertext string) string {
		aes, err := aes.NewCipher([]byte(salt))
		if err !=nil {
			panic(err)
		}
		gcm, err := cipher.NewGCM(aes)
		if err !=nil {
			panic(err)
		}
	// Since we know the ciphertext is actually nonce + ciphertext
	// and len(nonce) == NonceSize() we can separate the two using:
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err !=nil {
		panic(err)
	}
	return string(plaintext)
	}

	func main() {
		seed_String := make([]byte, 32)
		_, err := rand.Read(seed_String)
		if err !=nil {
			fmt.Print(err)
		}

		// Print our new random salt string generated from our original 32byte seed string
		fmt.Printf("\nRandom_Salt_String: %x\n\n", seed_String)

		// this will successfully encrypt / decrypt
		ciphertext1 := encrypt("Super (Secret Squirrel) info")
		fmt.Printf("Encrypted ciphertext 1: %x \n", ciphertext1 )

		plaintext1 := decrypt(ciphertext1)
		fmt.Printf("Decrypted plaintext 1: %s \n", plaintext1)

		// Without gcm + aes the (Hello <16bytes") string below would not satisfy the 16 bytes long min block size used by aes
		// and would cause the program to panic
		ciphertext2 := encrypt ("Hello <16bytes")
		fmt.Printf("\nEncypted ciphertext 2: %x \n", ciphertext2)

		plaintext2 := decrypt(ciphertext2)
		fmt.Printf("Decrypted plaintext 2: %s \n", plaintext2)

	fmt.Println("\nNote: Strings < 16bytes would normally cause the program to panic if using AES alone...")
	fmt.Println("Using AES + GCM aleviates the minimum size limitation (16bytes block size) and...")
	fmt.Println("also gives us message authentication (integrity) = Authenticated Encryption\n")

	fmt.Println("\nRunning the program in multiple terminal windows shows a unique salt + encryption hashes")
	fmt.Println("for each instance of the program that you initiate\n")

	// improved version of: dev.to/breda/secret-key-encryption-with-go-using-aes-316d

	}

