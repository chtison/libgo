package cli

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

// Template ...
const Template = `Usage:  {{.Path}} {{with .Arguments}}{{.}}{{else}}COMMAND{{end}}
{{with .Synopsys}}
{{.}}{{end}}
{{- with .Options}}

Options:
{{.}}{{end}}
{{- with .Commands}}
Commands:
{{.}}{{end}}
{{- with .Footer}}
{{.}}
{{end}}`

// Usage ...
func Usage(cmd *Command, args ...string) error {
	t := template.Must(template.New("usage").Parse(Template))
	path := usageFormatPath(cmd)
	m := map[string]interface{}{
		"Path":      path,
		"Arguments": cmd.Usage.Arguments,
		"Synopsys":  cmd.Usage.Synopsys,
		"Options":   usageFormatOptions(cmd.Flags),
		"Commands":  usageFormatCommands(cmd.Children),
		"Footer":    usageFormatFooter(cmd.Children, path),
	}
	t.Execute(os.Stderr, m)
	return nil
}

func usageFormatPath(cmd *Command) string {
	slice := make([]string, 0, 1)
	for p := cmd; p != nil; p = p.parent {
		slice = append(slice, p.name)
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return strings.Join(slice, " ")
}

func usageFormatOptions(set *FlagSet) string {
	if set == nil {
		return ""
	}
	keys := make([]string, 0, len(set.long))
	for k := range set.long {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		flag := set.long[k]
		buf.WriteString("  ")
		if shortName := flag.ShortName(); shortName != 0 {
			buf.WriteString("-" + string(shortName) + ", ")
		} else {
			buf.WriteString("    ")
		}
		longName := flag.LongName()
		flagType := flag.Type()
		buf.WriteString("--" + longName + " " + flagType)
		for i := 18 - len(longName) - len(flagType); i > 0; i-- {
			buf.WriteRune(' ')
		}
		buf.WriteString(flag.Usage().Synopsys)
		buf.WriteString(fmt.Sprintln(""))
	}
	return buf.String()
}

func usageFormatCommands(set CommandSet) string {
	if len(set) == 0 {
		return ""
	}
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	var buf bytes.Buffer
	for _, k := range keys {
		cmd := set[k]
		buf.WriteString("  " + cmd.name)
		for i := 14 - len(cmd.name); i > 0; i-- {
			buf.WriteRune(' ')
		}
		buf.WriteString(cmd.Usage.Synopsys + "\n")
	}
	return buf.String()
}

func usageFormatFooter(set CommandSet, path string) string {
	if len(set) > 0 {
		return fmt.Sprintf(`Run '%s COMMAND --help' for more information on a command.`, path)
	}
	return ""
}
