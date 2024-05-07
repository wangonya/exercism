package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	app_map := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for len(log) > 0 {
		r, size := utf8.DecodeRuneInString(log)
		if app, exists := app_map[r]; exists {
			return app
		}
		log = log[size:]
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog []byte
	for _, char := range log {
		if char == oldRune {
			newLog = utf8.AppendRune(newLog, newRune)
		} else {
			newLog = utf8.AppendRune(newLog, char)
		}
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
