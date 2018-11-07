package errors

import (
  "errors"
  "fmt"
  "strings"
)

type Error []error

func (err Error) Error() string {
  var b strings.Builder
  for _, err := range err {
    err := err.Error()
    if strings.HasSuffix(err, "\n") {
      fmt.Fprint(&b, err)
    } else {
      fmt.Fprintln(&b, err)
    }
  }
  return b.String()
}

func New(text string) Error {
  return []error{errors.New(text)}
}
