#!/bin/sh -e --

cd -P -- `dirname -- $0`

FMTDOC=`godoc fmt`
FUNCTIONS=`echo "$FMTDOC" | awk -F'[ (]+' '/^func /{print $2}'`
TYPES=`echo "$FMTDOC" | awk -F'[ ]+' '/^type /{print $2}'`

cat > fmt.go << EOF
package fmt
import "fmt"
// ...
var (
`for F in $FUNCTIONS; do echo "\t$F = fmt.$F"; done`
)
type (
`for T in $TYPES; do echo "\t// $T ...\n\t$T = fmt.$T"; done`
)
EOF

gofmt -w fmt.go

echo "`dirname -- $0`/fmt.go successfully created"
