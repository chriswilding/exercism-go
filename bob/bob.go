package bob

import "strings"

// Hey returns Bob's response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	question := strings.HasSuffix(remark, "?")
	yell := remark != strings.ToLower(remark) && remark == strings.ToUpper(remark)

	if question && yell {
		return "Calm down, I know what I'm doing!"
	}
	if question {
		return "Sure."
	}
	if yell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
