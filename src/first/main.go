package main

import (
  "fmt"
  "os"
  "./filter"
  "errors"
)

type person struct {
  name string
  age string
}


func main() {
  var name string
  var age string

  args := os.Args

  if len(args) < 3 {
    err := errors.New("Not enough arguments! Please provide an age and a name.")
    fmt.Println(err)
    os.Exit(1)
  }


  filteredArgs, err := filter.FilterArgs(args)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, argument := range(filteredArgs) {
    switch argument.Key {
      case "name":
        name = argument.Value
      case "age":
        age = argument.Value
    }
  }

  person := person{name: name, age: age}

  welcomeMessage := fmt.Sprintf("Welcome, %s! You're %s years old.", person.name, person.age)

  fmt.Println(welcomeMessage)

}


