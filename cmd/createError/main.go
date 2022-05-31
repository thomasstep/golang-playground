package main

import (
  "errors"
  "log"
)

func createError() error {
  return errors.New("This is my error")
}

func main() {
  err := createError()
  if err != nil {
    log.Fatal(err)
  }
}
