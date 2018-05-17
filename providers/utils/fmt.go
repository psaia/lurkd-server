package utils

import "regexp"

// FormatVersion will strip a version down to its bare numbers and decimals.
func FormatVersion(v string) string {
	re := regexp.MustCompile("[^0-9.]")
	return re.ReplaceAllString(v, "")
}
