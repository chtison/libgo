package flag

import "strings"

// Flag ...
type Flag interface {
	Name() (shortName, longName string)
	Usage() string
	Parse(value *string) error
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

func (flag *flag) errFlagNeedsValue() ErrFlagNeedsValue {
	if flag.longName != "" {
		return ErrFlagNeedsValue("--" + flag.longName)
	}
	return ErrFlagNeedsValue("-" + flag.shortName)
}

// FlagSet ...
type FlagSet struct {
	flags []Flag
}

// NewFlagSet ...
func NewFlagSet() *FlagSet {
	return &FlagSet{
		flags: make([]Flag, 0, 1),
	}
}

func (set *FlagSet) find(shortName, longName string) (index int, flag Flag) {
	for index, flag = range set.flags {
		shortFlagName, longFlagName := flag.Name()
		if (shortName == "" || shortName == shortFlagName) && (longName == "" || longName == longFlagName) {
			return index, flag
		}
	}
	return -1, nil
}

// Find ...
func (set *FlagSet) Find(shortName, longName string) Flag {
	_, flag := set.find(shortName, longName)
	return flag
}

// Remove ...
func (set *FlagSet) Remove(shortName, longName string) {
	index, flag := set.find(shortName, longName)
	if flag != nil {
		set.flags = append(set.flags[:index], set.flags[index+1:]...)
	}
}

// Add ...
func (set *FlagSet) Add(flag Flag) {
	set.Remove(flag.Name())
	set.flags = append(set.flags, flag)
}
