package cli

import "fmt"

type (
	// Flag ...
	Flag interface {
		ShortName() rune
		LongName() string
		Type() string
		Parse(arg *string) error
		Usage() *FlagUsage
	}
	// FlagUsage ...
	FlagUsage struct {
		Synopsys string
	}
	// FlagSet ...
	FlagSet struct {
		short map[rune]Flag
		long  map[string]Flag
	}
	// flag ...
	flag struct {
		shortName rune
		longName  string
		usage     FlagUsage
	}
)

////
//////// flag
////

func newFlag(shortName rune, longName string) *flag {
	if longName == "" {
		panic("A flag can't have an empty long name")
	}
	if shortName == '-' {
		panic("A flag can't have '-' as short name")
	}
	return &flag{
		shortName: shortName,
		longName:  longName,
	}
}

func (flag *flag) ShortName() rune {
	return flag.shortName
}

func (flag *flag) LongName() string {
	return flag.longName
}

func (flag *flag) Usage() *FlagUsage {
	return &flag.usage
}

////
//////// FlagSet
////

// NewFlagSet ...
func NewFlagSet() *FlagSet {
	return &FlagSet{
		short: make(map[rune]Flag, 0),
		long:  make(map[string]Flag, 0),
	}
}

// Add ...
func (set *FlagSet) Add(flags ...Flag) {
	for _, flag := range flags {
		if flag == nil {
			continue
		}
		shortName, longName := flag.ShortName(), flag.LongName()
		if shortName != 0 {
			if set.short[shortName] != nil {
				panic(fmt.Sprintf(`A flag -%c already exists`, shortName))
			}
		}
		if longName != "" {
			if set.long[longName] != nil {
				panic(fmt.Sprintf(`A flag --%s already exists`, longName))
			}
			set.long[longName] = flag
			if shortName != 0 {
				set.short[shortName] = flag
			}
		}
	}
}

// Find ...
func (set *FlagSet) Find(shortName rune, longName string) Flag {
	var f1, f2 Flag
	if shortName != 0 {
		f1 = set.short[shortName]
	}
	if longName != "" {
		f2 = set.long[longName]
	}
	if f1 == f2 || f2 == nil {
		return f1
	}
	if f1 == nil {
		return f2
	}
	return nil
}
