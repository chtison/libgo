package generated

import (
	"strings"
	"unicode"
)

type Strings struct{}

func NewStrings() *Strings { return &Strings{} }

func (_ *Strings) Compare(a string, b string) int {
	return strings.Compare(a, b)
}

func (_ *Strings) Contains(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func (_ *Strings) ContainsAny(s string, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func (_ *Strings) ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func (_ *Strings) Count(s string, substr string) int {
	return strings.Count(s, substr)
}

func (_ *Strings) EqualFold(s string, t string) bool {
	return strings.EqualFold(s, t)
}

func (_ *Strings) Fields(s string) []string {
	return strings.Fields(s)
}

func (_ *Strings) FieldsFunc(s string, f func(rune) bool) []string {
	return strings.FieldsFunc(s, f)
}

func (_ *Strings) HasPrefix(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func (_ *Strings) HasSuffix(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func (_ *Strings) Index(s string, substr string) int {
	return strings.Index(s, substr)
}

func (_ *Strings) IndexAny(s string, chars string) int {
	return strings.IndexAny(s, chars)
}

func (_ *Strings) IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

func (_ *Strings) IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

func (_ *Strings) IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

func (_ *Strings) Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func (_ *Strings) LastIndex(s string, substr string) int {
	return strings.LastIndex(s, substr)
}

func (_ *Strings) LastIndexAny(s string, chars string) int {
	return strings.LastIndexAny(s, chars)
}

func (_ *Strings) LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}

func (_ *Strings) LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

func (_ *Strings) Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

func (_ *Strings) NewReader(s string) *strings.Reader {
	return strings.NewReader(s)
}

func (_ *Strings) NewReplacer(oldnew ...string) *strings.Replacer {
	return strings.NewReplacer(oldnew...)
}

func (_ *Strings) Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func (_ *Strings) Replace(s string, old string, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func (_ *Strings) Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func (_ *Strings) SplitAfter(s string, sep string) []string {
	return strings.SplitAfter(s, sep)
}

func (_ *Strings) SplitAfterN(s string, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

func (_ *Strings) SplitN(s string, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

func (_ *Strings) Title(s string) string {
	return strings.Title(s)
}

func (_ *Strings) ToLower(s string) string {
	return strings.ToLower(s)
}

func (_ *Strings) ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

func (_ *Strings) ToTitle(s string) string {
	return strings.ToTitle(s)
}

func (_ *Strings) ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

func (_ *Strings) ToUpper(s string) string {
	return strings.ToUpper(s)
}

func (_ *Strings) ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

func (_ *Strings) Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

func (_ *Strings) TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

func (_ *Strings) TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

func (_ *Strings) TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

func (_ *Strings) TrimPrefix(s string, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func (_ *Strings) TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func (_ *Strings) TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

func (_ *Strings) TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func (_ *Strings) TrimSuffix(s string, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
