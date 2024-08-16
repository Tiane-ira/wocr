package utils

import "regexp"

var (
	regVin = regexp.MustCompile(`[A-HJ-NPR-XZ\d]{8}[X\d][A-HJ-NPR-Z\d]{3}\d{5}`)
)

func GetVincode(text string) (bool, string) {
	matcher := regVin.FindAllString(text, -1)
	if len(matcher) > 0 {
		return true, matcher[0]
	}
	return false, ""
}
