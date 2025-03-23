package cmdargs

import "strings"

func ParseArguments(input string) (path string, searchQuery []string) {
	segments := strings.Split(input, "~")

	if len(segments) > 0 {
		path = segments[0]
	}

	if len(segments) > 1 {
		searchQuery = strings.Split(segments[1], "+")
	}

	return
}
