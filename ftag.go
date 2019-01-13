package ftag

import (
	"strconv"
	"strings"
)

// Tag represents the pre-parsed tag string of a Go struct field.
type Tag map[string][]string

func (t Tag) Contains(key, value string) bool {
	for _, v := range t[key] {
		if v == value {
			return true
		}
	}
	return false
}

func (t Tag) HasOption(key, value string) bool {
	if len(t[key]) > 0 {
		for _, v := range t[key][1:] {
			if v == value {
				return true
			}
		}
	}
	return false
}

func (t Tag) First(key string) string {
	if vv := t[key]; len(vv) > 0 {
		return vv[0]
	}
	return ""
}

func (t Tag) Second(key string) string {
	if vv := t[key]; len(vv) > 1 {
		return vv[1]
	}
	return ""
}

// New parses the given string and returns a new Tag value.
//
// NOTE: The loop's logic is a slightly modified copy of the
// StructTag.Lookup method from Go's reflect package.
func New(str string) Tag {
	if str == "" {
		return nil
	}

	tag := make(Tag)
	for str != "" {
		i := 0
		for i < len(str) && str[i] == ' ' {
			i++
		}
		str = str[i:]
		if str == "" {
			break
		}

		i = 0
		for i < len(str) && str[i] > ' ' && str[i] != ':' && str[i] != '"' && str[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(str) || str[i] != ':' || str[i+1] != '"' {
			break
		}
		key := string(str[:i])
		str = str[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(str) && str[i] != '"' {
			if str[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(str) {
			break
		}
		qvalue := string(str[:i+1])
		str = str[i+1:]

		value, err := strconv.Unquote(qvalue)
		if err != nil {
			break
		}
		tag[key] = strings.Split(value, ",")
	}
	return tag
}