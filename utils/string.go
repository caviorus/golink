package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// GenerateSlug generates a URL-friendly slug from input string
func GenerateSlug(input string) string {
	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace non-alphanumeric with dash
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	// Trim leading/trailing dashes
	slug = strings.Trim(slug, "-")

	return slug
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
