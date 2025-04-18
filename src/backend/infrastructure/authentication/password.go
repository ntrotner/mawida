package authentication

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes a password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// checkPasswordHash checks if a password matches a hash
func checkPasswordHash(password *string, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}

// createPasswordSalt generates a pseudo-random salt
func createPasswordSalt(length int) string {
	accumulatedSalt := []string{}

	for len(accumulatedSalt) <= length {
		r := rand.New(rand.NewSource(time.Now().Local().UnixMilli() + int64(len(accumulatedSalt))))
		salt := string(byte(r.Uint32()))

		accumulatedSalt = append(accumulatedSalt, salt)
	}

	return strings.Join(accumulatedSalt, "")
}

// CreatePassword creates a password
func CreatePassword(password *string) (string, string, error) {
	salt := createPasswordSalt(20)
	hash, err := hashPassword(*password + salt)

	return hash, salt, err
}

// CheckPassword checks if a password matches a hash
func CheckPassword(password *string, hash *string, salt *string) bool {
	saltedPassword := *password + *salt
	return checkPasswordHash(&saltedPassword, hash)
}
