package security

import (
	"git.nextchat.org/nextchat/nextchat-go/utils"
	"github.com/matthewhartstonge/argon2"
)

func EncryptArgon(text string) (string, bool) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(text))
	utils.CheckError(err)
	return string(encoded), true
}

func CheckEncript(password, text string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(text))
	utils.CheckError(err)
	return ok
}
