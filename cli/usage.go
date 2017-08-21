package cli

import (
	"bytes"
	"os"
	"sort"
	"text/template"
)

// Template ...
const Template = `Usage:  {{.Name}} {{with .Arguments}}{{.}}{{else}}COMMAND{{end}}
{{- with .Synopsys}}

{{.}}{{end}}
{{- with .Options}}

Options:
{{.}}{{end}}
{{- with .Commands}}

Commands:
{{.}}{{end}}
{{- with .Footer}}

{{.}}{{end}}
`

// Usage ...
func Usage(cmd *Command, args ...string) error {
	t := template.Must(template.New("usage").Parse(Template))
	m := map[string]interface{}{
		"Name":      cmd.name,
		"Arguments": cmd.Usage.Arguments,
		"Synopsys":  cmd.Usage.Synopsys,
		"Options":   usageFormatOptions(cmd.Flags),
		"Commands":  usageFormatCommands(cmd.children),
		"Footer":    usageFormatFooter(cmd.Flags),
	}
	t.Execute(os.Stderr, m)
	return nil
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
			buf.WriteByte(' ')
		}
		buf.WriteString(flag.Usage().Synopsys + "\n")
	}
	return buf.String()
}

func usageFormatCommands(set CommandSet) string {
	return ""
}

func usageFormatFooter(set *FlagSet) string {
	return ""
}
