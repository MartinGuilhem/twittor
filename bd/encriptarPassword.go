package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword is the routine that encripts a password */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err

}
