package hoshinova

import (
	"unicode/utf8"
)

// https://dev.to/takakd/go-safe-truncate-string-9h0#comment-1hp14
func TruncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}

	if utf8.RuneCountInString(str) < length {
		return str
	}

	return string([]rune(str)[:length])
}

func EscapeStatus(status interface{}) string {
	switch v := status.(type) {
	case map[string]interface{}:
		return "waiting"
	case string:
		return v
	default:
		return "unknown"
	}
}
