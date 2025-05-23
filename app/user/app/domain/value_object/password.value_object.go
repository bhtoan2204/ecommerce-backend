package value_object

import (
	"regexp"
	"strings"
)

type Password string

func (p Password) String() string {
	return string(p)
}

func (p Password) IsEmpty() bool {
	return p == ""
}

func (p Password) IsValid() bool {
	pass := string(p)
	if len(pass) < 8 {
		return false
	}
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(pass)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(pass)
	return hasLetter && hasDigit
}

func (p Password) Mask() string {
	pass := string(p)
	if len(pass) <= 2 {
		return strings.Repeat("*", len(pass))
	}
	return pass[:1] + strings.Repeat("*", len(pass)-2) + pass[len(pass)-1:]
}
