// Package tagutil provides a couple helper methods for dealing with Go's struct tags.
//
// This package follows the convention outlined in the Go's reflect package
// documentation at: https://golang.org/pkg/reflect/#StructTag, ie By convention,
// tag strings are a concatenation of optionally space-separated key:"value" pairs. [...]
//
// Additionally the "value" can be a comma-separated list of items, in which case
// the first item is regarded as the "main" value and all of the subsequent items
// are considered as "options" of that pair.
package tagutil

import (
	"strconv"
	"strings"
)

// Tag represents a Go struct tag.
type Tag map[string][]string

// Get returns the value associated with key in the Tag. If there is no such
// key in the tag, Get returns the empty string. If the tag does not have the
// conventional format, the value returned by Get is unspecified.
func (t Tag) Get(key string) string {
	return strings.Join(t[key], ",")
}

func (t Tag) Len(key string) int {
	return len(t[key])
}

// Contains reports whether the value associated with the given key matches
// the provided string val.
func (t Tag) Contains(key, val string) bool {
	for _, v := range t[key] {
		if v == val {
			return true
		}
	}
	return false
}

// HasOption reports whether at least one "option" that is associated with
// the given key matches the provided string val.
func (t Tag) HasOption(key, val string) bool {
	if len(t[key]) > 0 {
		for _, v := range t[key][1:] {
			if v == val {
				return true
			}
		}
	}
	return false
}

func (t Tag) NumOptions(key string) int {
	if len(t[key]) > 0 {
		return len(t[key][1:])
	}
	return 0
}

// First returns the "main" value associated with key in the Tag.
// If there is no such key in the tag, First returns the empty string.
func (t Tag) First(key string) string {
	if vv := t[key]; len(vv) > 0 {
		return vv[0]
	}
	return ""
}

// Second returns the first "option" value associated with key in the Tag.
// If there is no such key in the tag, or if there are no options associated
// with that key, Second returns the empty string.
func (t Tag) Second(key string) string {
	if vv := t[key]; len(vv) > 1 {
		return vv[1]
	}
	return ""
}

// New parses the given string and returns a new Tag value.
func New(str string) Tag {
	if str == "" {
		return nil
	}

	tag := make(Tag)

	// NOTE: The loop's logic is a slightly modified copy of the
	// StructTag.Lookup method from Go's reflect package.
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