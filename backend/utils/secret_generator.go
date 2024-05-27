package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomToken(length int) string {
	arr := make([]byte, length)
	_, err := rand.Read(arr)
	if err != nil {
		fmt.Println("Error")
	}
	token := base64.URLEncoding.EncodeToString(arr)
	return token
}
