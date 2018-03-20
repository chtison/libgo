package generated

import (
	"strings"
	"unicode"
)

type Strings struct{}

func NewStrings() *Strings { return &Strings{} }

func (*Strings) Compare(a string, b string) int {
	return strings.Compare(a, b)
}

func (*Strings) Contains(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func (*Strings) ContainsAny(s string, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func (*Strings) ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func (*Strings) Count(s string, substr string) int {
	return strings.Count(s, substr)
}

func (*Strings) EqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func (*Strings) Fields(s string) []string {
	return strings.Fields(s)
}

func (*Strings) FieldsFunc(s string, f func(rune) bool) []string {
	return strings.FieldsFunc(s, f)
}

func (*Strings) HasPrefix(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func (*Strings) HasSuffix(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func (*Strings) Index(s string, substr string) int {
	return strings.Index(s, substr)
}

func (*Strings) IndexAny(s string, chars string) int {
	return strings.IndexAny(s, chars)
}

func (*Strings) IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

func (*Strings) IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

func (*Strings) IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

func (*Strings) Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func (*Strings) LastIndex(s string, substr string) int {
	return strings.LastIndex(s, substr)
}

func (*Strings) LastIndexAny(s string, chars string) int {
	return strings.LastIndexAny(s, chars)
}

func (*Strings) LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}

func (*Strings) LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

func (*Strings) Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

func (*Strings) NewReader(s string) *strings.Reader {
	return strings.NewReader(s)
}

func (*Strings) NewReplacer(oldnew ...string) *strings.Replacer {
	return strings.NewReplacer(oldnew...)
}

func (*Strings) Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func (*Strings) Replace(s string, old string, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func (*Strings) Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func (*Strings) SplitAfter(s string, sep string) []string {
	return strings.SplitAfter(s, sep)
}

func (*Strings) SplitAfterN(s string, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

func (*Strings) SplitN(s string, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

func (*Strings) Title(s string) string {
	return strings.Title(s)
}

func (*Strings) ToLower(s string) string {
	return strings.ToLower(s)
}

func (*Strings) ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

func (*Strings) ToTitle(s string) string {
	return strings.ToTitle(s)
}

func (*Strings) ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

func (*Strings) ToUpper(s string) string {
	return strings.ToUpper(s)
}

func (*Strings) ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

func (*Strings) Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

func (*Strings) TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

func (*Strings) TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func (*Strings) TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

func (*Strings) TrimPrefix(s string, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func (*Strings) TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func (*Strings) TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

func (*Strings) TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func (*Strings) TrimSuffix(s string, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
