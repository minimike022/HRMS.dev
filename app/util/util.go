package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) string {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed_password)
}

func CompareHash(stored_password string, forms_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(stored_password), []byte(forms_password))
	return err
}

func EncryptToken() {

}


func DecryptToken() {

}



