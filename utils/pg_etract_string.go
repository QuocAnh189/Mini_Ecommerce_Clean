package utils

import "regexp"

func ExtractConstraintName(err error) string {
	if err == nil {
		return ""
	}

	re := regexp.MustCompile(`unique constraint "(.*?)"`)
	match := re.FindStringSubmatch(err.Error())

	if len(match) > 1 {
		return match[1]
	}
	return "unique_violation"
}
