package helper

import (
	"strings"
	"crypto/rand"
	"encoding/hex"
	)
	
func GRandomString() string  {
	id := make([]byte, 8)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
	}
	return "3EB0" + strings.ToUpper(hex.EncodeToString(id)) + "BALL"
}

