package gocommon

import "strings"

type String string

func (s String) IsEmpty() bool {
	return len(s) == 0
}

func (s String) Length() int {
	return len(s)
}

func (s String) toLower() String {
	return String(strings.ToLower(string(s)))
}

func (s String) toUpper() String {
	return String(strings.ToUpper(string(s)))
}

func (s String) Contains(sub string) bool {
	return strings.Contains(string(s), sub)
}

// Replace all instances of old substring with new value.
// Return new String
func (s String) Replace(oldStr string, newStr string) String {
	return String(strings.Replace(string(s), oldStr, newStr, -1))
}

func (s String) Split(keyword string) []String {
	s1 := strings.Split(string(s), keyword)
	var s2 []String
	for _, e := range s1 {
		s2 = append(s2, String(e))
	}
	return s2
}

func (s String) Index(keyword string) int {
	return strings.Index(string(s), keyword)
}
