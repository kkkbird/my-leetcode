package isMatch

import (
	"regexp"
)

const (
	TypeExact = 0
	TypeAny   = 1 // *
	TypeOnce  = 2 // ? no need to support
)

type Pattern struct {
	c rune
	t int
}

func parsePattern(p string) []Pattern {
	pp := make([]Pattern, 0)
	for _, c := range p {
		if (c >= 'a' && c <= 'z') || c == '.' {
			pp = append(pp, Pattern{c: c, t: TypeExact})
		} else if c == '*' {
			pp[len(pp)-1].t = TypeAny
		}
	}

	return pp
}

func recursiveMatch(s string, pp []Pattern) bool {
	if len(pp) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}

	pp1 := pp[0]
	switch pp1.t {
	case TypeAny:
		if len(s) == 0 {
			return recursiveMatch("", pp[1:])
		}

		c := rune(s[0])
		if pp1.c != '.' && c != pp1.c {
			return recursiveMatch(s, pp[1:])
		}

		// try not use * pattern
		if recursiveMatch(s, pp[1:]) {
			return true
		}

		// use * pattern
		return recursiveMatch(s[1:], pp)

	case TypeExact:
		if len(s) == 0 {
			return false
		}
		c := rune(s[0])
		if pp1.c != '.' && c != pp1.c {
			return false
		}
		return recursiveMatch(s[1:], pp[1:])
	}

	return false
}

func isMatch(s string, p string) bool {
	if p == "" {
		if s == "" {
			return true
		} else {
			return false
		}
	}

	pp := parsePattern(p)

	return recursiveMatch(s, pp)
}

func isMatch2(s string, p string) bool {
	reg := regexp.MustCompile("^" + p + "$")
	return reg.Match([]byte(s))
}
