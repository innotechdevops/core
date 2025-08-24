package core

import (
	"fmt"
	"net/url"
	"strings"
)

func ToQuery(params map[string]string) string {
	var sb strings.Builder
	for key, value := range params {
		if sb.Len() > 0 {
			sb.WriteString("&") // Add '&' between parameters
		}
		sb.WriteString(key + "=" + value) // Append key=value
	}
	return sb.String()
}

func Query(rawURL string) url.Values {
	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	// Get the query parameters
	queryParams := parsedURL.Query()

	return queryParams
}
