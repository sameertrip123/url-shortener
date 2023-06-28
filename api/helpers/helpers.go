package helpers

import (
	"os"
	"strings"
)

func EnforeHTTP(url string) string {
	// Check if the first 4 character of the url are http or not, if they are not then add http:// + url
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMIAN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL == os.Getenv("DOMIAN")
}
