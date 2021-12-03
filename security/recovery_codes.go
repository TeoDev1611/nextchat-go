package security

import (
	"git.nextchat.org/nextchat/nextchat-go/utils"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func genId() string {
	id, err := nanoid.New(16)
	utils.CheckError(err)
	return id
}

func GenerateCodes() []string {
	info := []string{genId(), genId(), genId(), genId(), genId(), genId()}
	return info
}
