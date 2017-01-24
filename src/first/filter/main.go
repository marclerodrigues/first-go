package filter

import (
  "errors"
  "strings"
)

type arg struct {
  Key string
  Value string
}


func FilterArgs(args []string) ([]arg, error) {
  var filteredArgs []arg

  if len(args) <= 1 {
    err := errors.New("No args were passed!")
    return filteredArgs, err
  }

  for index, argument := range(args) {
    if index != 0 {
      a := strings.Split(argument, "=")
      new_arg := arg{Key: a[0], Value: a[1]}
      filteredArgs = append(filteredArgs, new_arg)
    }
  }

  return filteredArgs, nil
}
