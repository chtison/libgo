#!/bin/sh -e --

PACKAGES="
			fmt
			strings
			time
"

cd -P -- `dirname -- "$0"`

DIR=generated

mkdir -p "$DIR"

generate() {

PACKAGE=`basename "$1"`
FILE="$DIR/$PACKAGE.go"

FUNCTIONS=`godoc "$1" | grep '^func [^(]' | cut -d' ' -f2-`
TYPE=`echo "$PACKAGE" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}'`

cat > $FILE << EOF
package generated
import "$1"
// $TYPE ...
type $TYPE struct {
$(echo "$FUNCTIONS" | while read -r F; do
	NAME=`echo $F | cut -d'(' -f1`
	SIGN=`echo $F | cut -d'(' -f2- | sed -E "s/([ \*(])([[:upper:]])/\1$PACKAGE.\2/g"`
	echo "$NAME" 'func(' "$SIGN"
done)
}
// New$TYPE ...
func New$TYPE() *$TYPE {
	return &$TYPE{
		$(echo "$FUNCTIONS" | while read -r F; do
			NAME=`echo $F | cut -d'(' -f1`
			echo "$NAME" ':' "$PACKAGE" '.' "$NAME" ','
		done)
	}
}
EOF

goimports -w "$FILE"

echo "$FILE successfully created"

}

for P in $PACKAGES; do
	generate $P
done
