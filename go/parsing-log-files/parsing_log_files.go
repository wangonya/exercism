package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`^(\[TRC]|\[DBG]|\[INF]|\[WRN]|\[ERR]|\[FTL])`)
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<(~|-|=|\*)*>`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	r := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, s := range lines {
		if r.FindString(s) != "" {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line(\d+)?`)
	return r.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile(`User +(\w*)`)
	newLines := []string{}
	for _, line := range lines {
		match := r.FindString(line)
		if match != "" {
			newLines = append(newLines, "[USR] "+strings.Fields(match)[1]+" "+line)
		} else {
			newLines = append(newLines, line)
		}
	}
	return newLines
}
