package helpers

import (
	"fmt"
	"os"
)

func GetPublicPath(key string) string {
	hostAddress := os.Getenv("APPLICATION_HOST")
	publicPathURL := fmt.Sprintf("%v/api/attachment/%v", hostAddress, key)
	return publicPathURL
}
