package utils

import (
	"regexp"
	"strings"
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
