package utils

import (
	"fmt"
	"os"
)

func CreateImageLink(path string) string {
	var host = os.Getenv("HOST")
	if path == "" {
		return ""
	}
	return fmt.Sprintf("%s/assets/images/items/%s%s", host, path, "?time="+GetCurrentDateTimeString())
}

func CreateImageLinks(images []string) []string {
	result := make([]string, len(images))
	for i, img := range images {
		result[i] = CreateImageLink(img)
	}
	return result
}
