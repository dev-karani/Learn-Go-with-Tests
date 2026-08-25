package iteration

import "strings"

const repeatCount = 5

func Repeat(charachter string) string {
	var repeated strings.Builder

	for range repeatCount {
		repeated.WriteString(charachter)
	}
	return repeated.String()
}
