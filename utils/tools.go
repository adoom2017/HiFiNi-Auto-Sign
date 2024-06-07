package utils

import "regexp"

func regexpMatch(matched string, pattern string) string {
	r, _ := regexp.Compile(pattern)

	ret := r.FindStringSubmatch(matched)
	if ret == nil {
		return ""
	} else {
		return ret[1]
	}
}

func GetSign(matched string) string {
	return regexpMatch(matched, "var sign = \"([^\"]+)\"")
}

func GetUsername(matched string) string {
	return regexpMatch(matched, "<li class=\"nav-item username\"><a class=\"nav-link\" href=\"my.htm\"><img class=\"avatar-1\" src=\".*?\"> (.*?)</a></li>")
}

type Message interface {
	SendMessage(text string)
}

func Sending(m Message, text string) {
	m.SendMessage(text)
}
