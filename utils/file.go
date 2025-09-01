package utils

import (
	"fmt"
	"os"
)

func CreateImageLink(path string) string {
	var host = os.Getenv("HOST")
	return fmt.Sprintf("%s/assets/images/items/%s", host, path)
}

func CreateImageLinks(images []string) []string {
	result := make([]string, len(images))
	for i, img := range images {
		result[i] = CreateImageLink(img)
	}
	return result
}
