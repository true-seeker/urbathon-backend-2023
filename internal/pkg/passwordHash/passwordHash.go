package passwordHash

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func GenerateRandomSalt() []byte {
	var salt = make([]byte, 16)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)

	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func DoPasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	fmt.Println(hashedPassword, currPasswordHash)
	return hashedPassword == currPasswordHash
}
