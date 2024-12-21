package name

import (
	"github.com/samber/lo"
	"strings"
)

type Name struct {
	words []string
}

func New(s string) Name {
	rs := []rune(strings.Join(strings.Fields(s), " "))
	ws := [][]rune{}
	for i, r := range rs {
		if i == 0 {
			ws = append(ws, []rune{r})
			continue
		}
		pr := rs[i-1]
		switch {
		case isUpperRune(r):
			if isLowerRune(pr) || isDigitRune(pr) || isSymbolRune(pr) {
				ws = append(ws, []rune{r})
			} else {
				ws[len(ws)-1] = append(ws[len(ws)-1], r)
			}
		case isLowerRune(r):
			if isSymbolRune(pr) {
				ws = append(ws, []rune{r})
			} else {
				ws[len(ws)-1] = append(ws[len(ws)-1], r)
			}
		case isDigitRune(r):
			if isSymbolRune(pr) {
				ws = append(ws, []rune{r})
			} else {
				ws[len(ws)-1] = append(ws[len(ws)-1], r)
			}
		default:
			ws = append(ws, []rune{r})
		}
	}
	return Name{
		words: lo.Map(ws, func(r []rune, _ int) string { return string(r) }),
	}
}

func (n Name) String() string {
	return n.Join("", "", "")
}

func (n Name) Equal(other Name) bool {
	return n.String() == other.String()
}

func (n Name) Join(sep, prefix, suffix string) string {
	s := lo.Map(n.words, func(w string, _ int) string { return w })
	return prefix + strings.Join(s, sep) + suffix
}

func (n Name) Map(f func(w string) string) Name {
	return Name{
		words: lo.Map(n.words, func(w string, _ int) string { return f(w) }),
	}
}

func (n Name) Append(s string) Name {
	return Name{
		words: append(append([]string{}, n.words...), New(s).words...),
	}
}

func (n Name) Prepend(s string) Name {
	return Name{
		words: append(New(s).words, n.words...),
	}
}

func (n Name) RemoveIf(f func(w string) bool) Name {
	return Name{
		words: lo.Filter(n.words, func(w string, _ int) bool { return !f(w) }),
	}
}

func (n Name) LowerCamel() string {
	u := n.UpperCamel()
	if u == "" {
		return ""
	}
	rs := []rune(u)
	return strings.ToLower(string(rs[0])) + string(rs[1:])
}

func (n Name) UpperCamel() string {
	return n.
		Map(ToFirstUpper).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("", "", "")
}

func (n Name) LowerSnake() string {
	return n.
		Map(ToLower).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("_", "", "")
}

func (n Name) AllUpperSnake() string {
	return n.
		Map(ToAllUpper).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("_", "", "")
}

func (n Name) FirstUpperSnake() string {
	return n.
		Map(ToAllUpper).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("_", "", "")
}

func (n Name) LowerKebab() string {
	return n.
		Map(ToLower).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("-", "", "")
}

func (n Name) FirstUpperKebab() string {
	return n.
		Map(ToFirstUpper).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("-", "", "")
}

func (n Name) AllUpperKebab() string {
	return n.
		Map(ToAllUpper).
		Map(ignoreSpecial).
		RemoveIf(isEmpty).
		Join("-", "", "")
}

func isLowerRune(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpperRune(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isDigitRune(r rune) bool {
	return '0' <= r && r <= '9'
}

func isSymbolRune(r rune) bool {
	return !(isLowerRune(r) || isUpperRune(r) || isDigitRune(r))
}

func ToLower(w string) string {
	return strings.ToLower(w)
}

func ToAllUpper(w string) string {
	return strings.ToUpper(w)
}

func ToFirstUpper(w string) string {
	rs := []rune(strings.ToLower(w))
	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}

func ignoreSpecial(w string) string {
	rs := []rune(w)
	if len(rs) == 1 && isSymbolRune(rs[0]) {
		return ""
	}
	return w
}

func isEmpty(w string) bool {
	return w == ""
}