// internal/util/password.go
package util

import "golang.org/x/crypto/bcrypt"

// HashPassword は生のパスワードからbcryptハッシュを生成するのだ
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash はハッシュと生のパスワードが一致するかを検証するのだ
// この関数も一緒に作っておくと便利なのだ
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
