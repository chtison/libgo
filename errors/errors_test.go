package errors

import (
  "errors"
  "testing"
)

func TestError(t *testing.T) {
  err := New("FAILED")
  if err.Error() != "FAILED\n" {
    t.Errorf("%q", err.Error())
  }
  err = append(err, errors.New("OMG"))
  if err.Error() != "FAILED\nOMG\n" {
    t.Errorf("%q", err.Error())
  }
  err = append(err, New("WOW"))
  if err.Error() != "FAILED\nOMG\nWOW\n" {
    t.Errorf("%q", err.Error())
  }
}
