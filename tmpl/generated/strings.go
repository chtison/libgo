package generated

import (
	"strings"
	"unicode"
)

// Strings ...
type Strings struct {
	Compare        func(a, b string) int
	Contains       func(s, substr string) bool
	ContainsAny    func(s, chars string) bool
	ContainsRune   func(s string, r rune) bool
	Count          func(s, substr string) int
	EqualFold      func(s, t string) bool
	Fields         func(s string) []string
	FieldsFunc     func(s string, f func(rune) bool) []string
	HasPrefix      func(s, prefix string) bool
	HasSuffix      func(s, suffix string) bool
	Index          func(s, substr string) int
	IndexAny       func(s, chars string) int
	IndexByte      func(s string, c byte) int
	IndexFunc      func(s string, f func(rune) bool) int
	IndexRune      func(s string, r rune) int
	Join           func(a []string, sep string) string
	LastIndex      func(s, substr string) int
	LastIndexAny   func(s, chars string) int
	LastIndexByte  func(s string, c byte) int
	LastIndexFunc  func(s string, f func(rune) bool) int
	Map            func(mapping func(rune) rune, s string) string
	Repeat         func(s string, count int) string
	Replace        func(s, old, new string, n int) string
	Split          func(s, sep string) []string
	SplitAfter     func(s, sep string) []string
	SplitAfterN    func(s, sep string, n int) []string
	SplitN         func(s, sep string, n int) []string
	Title          func(s string) string
	ToLower        func(s string) string
	ToLowerSpecial func(c unicode.SpecialCase, s string) string
	ToTitle        func(s string) string
	ToTitleSpecial func(c unicode.SpecialCase, s string) string
	ToUpper        func(s string) string
	ToUpperSpecial func(c unicode.SpecialCase, s string) string
	Trim           func(s string, cutset string) string
	TrimFunc       func(s string, f func(rune) bool) string
	TrimLeft       func(s string, cutset string) string
	TrimLeftFunc   func(s string, f func(rune) bool) string
	TrimPrefix     func(s, prefix string) string
	TrimRight      func(s string, cutset string) string
	TrimRightFunc  func(s string, f func(rune) bool) string
	TrimSpace      func(s string) string
	TrimSuffix     func(s, suffix string) string
	NewReader      func(s string) *strings.Reader
	NewReplacer    func(oldnew ...string) *strings.Replacer
}

// NewStrings ...
func NewStrings() *Strings {
	return &Strings{
		Compare:        strings.Compare,
		Contains:       strings.Contains,
		ContainsAny:    strings.ContainsAny,
		ContainsRune:   strings.ContainsRune,
		Count:          strings.Count,
		EqualFold:      strings.EqualFold,
		Fields:         strings.Fields,
		FieldsFunc:     strings.FieldsFunc,
		HasPrefix:      strings.HasPrefix,
		HasSuffix:      strings.HasSuffix,
		Index:          strings.Index,
		IndexAny:       strings.IndexAny,
		IndexByte:      strings.IndexByte,
		IndexFunc:      strings.IndexFunc,
		IndexRune:      strings.IndexRune,
		Join:           strings.Join,
		LastIndex:      strings.LastIndex,
		LastIndexAny:   strings.LastIndexAny,
		LastIndexByte:  strings.LastIndexByte,
		LastIndexFunc:  strings.LastIndexFunc,
		Map:            strings.Map,
		Repeat:         strings.Repeat,
		Replace:        strings.Replace,
		Split:          strings.Split,
		SplitAfter:     strings.SplitAfter,
		SplitAfterN:    strings.SplitAfterN,
		SplitN:         strings.SplitN,
		Title:          strings.Title,
		ToLower:        strings.ToLower,
		ToLowerSpecial: strings.ToLowerSpecial,
		ToTitle:        strings.ToTitle,
		ToTitleSpecial: strings.ToTitleSpecial,
		ToUpper:        strings.ToUpper,
		ToUpperSpecial: strings.ToUpperSpecial,
		Trim:           strings.Trim,
		TrimFunc:       strings.TrimFunc,
		TrimLeft:       strings.TrimLeft,
		TrimLeftFunc:   strings.TrimLeftFunc,
		TrimPrefix:     strings.TrimPrefix,
		TrimRight:      strings.TrimRight,
		TrimRightFunc:  strings.TrimRightFunc,
		TrimSpace:      strings.TrimSpace,
		TrimSuffix:     strings.TrimSuffix,
		NewReader:      strings.NewReader,
		NewReplacer:    strings.NewReplacer,
	}
}
