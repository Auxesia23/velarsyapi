package utils

import (
	"regexp"
	"strings"
)

// Slugify the given string
func ToSlug(title string) string {
	slug := strings.ToLower(title)

	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	slug = reg.ReplaceAllString(slug, "")

	reg2 := regexp.MustCompile(`-+`)
	slug = reg2.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
