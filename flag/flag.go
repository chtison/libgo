package flag

import "strings"

// Flag ...
type Flag interface {
	Name() (shortName, longName string)
	Usage() string
	ParseValue(value string) error
}

// flag ...
type flag struct {
	shortName string
	longName  string
	usage     string
}

// Newflag ...
func newFlag(shortName, longName string) *flag {
	if shortName == "" && longName == "" {
		panic("String can't have empty names.")
	}
	if len([]rune(shortName)) > 1 {
		panic("String's short name can't have more than one rune.")
	}
	if strings.Contains(longName, "=") {
		panic("String's long name cannot contains '='")
	}
	return &flag{
		shortName: shortName,
		longName:  longName,
	}
}

// Name ...
func (flag *flag) Name() (shortName, longName string) {
	return flag.shortName, flag.longName
}

// Usage ...
func (flag *flag) Usage() string {
	return flag.usage
}

// SetUsage ...
func (flag *flag) SetUsage(usage string) {
	flag.usage = usage
}

// Set ...
type Set struct {
	flags []Flag
}

// NewSet ...
func NewSet() *Set {
	return &Set{
		flags: make([]Flag, 0, 1),
	}
}

func (set *Set) find(shortName, longName string) (index int, flag Flag) {
	for index, flag = range set.flags {
		shortFlagName, longFlagName := flag.Name()
		if shortName != "" && shortName == shortFlagName || longName != "" && longName == longFlagName {
			return index, flag
		}
	}
	return -1, nil
}

// Find ...
func (set *Set) Find(shortName, longName string) Flag {
	_, flag := set.find(shortName, longName)
	return flag
}

func (set *Set) remove(atIndex int) {
	set.flags = append(set.flags[:atIndex], set.flags[atIndex+1:]...)
}

// Remove ...
func (set *Set) Remove(shortName, longName string) {
	index, flag := set.find(shortName, longName)
	if flag != nil {
		set.remove(index)
	}
}

// Add ...
func (set *Set) Add(flag Flag) {
	set.Remove(flag.Name())
	set.flags = append(set.flags, flag)
}
