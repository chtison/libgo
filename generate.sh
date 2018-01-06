#!/bin/sh -e --

(cd fmt && ./generate.sh)
(cd tmpl && ./generate.sh)
