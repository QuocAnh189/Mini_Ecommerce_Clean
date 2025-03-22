package utils

import "regexp"

func ExtractConstraintName(err error) string {
	if err == nil {
		return ""
	}

	// Biểu thức chính quy để tìm tên constraint
	re := regexp.MustCompile(`unique constraint "(.*?)"`)
	match := re.FindStringSubmatch(err.Error())

	if len(match) > 1 {
		return match[1] // Trả về tên constraint, ví dụ: "uni_users_email"
	}
	return "unique_violation"
}
