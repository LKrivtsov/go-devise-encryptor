# go-devise-encryptor
Devise encryptor written in golang. Useful for integrating a go application with a ruby on rails database. This is just a test to see how easy it is to integrate a go app with an existing rails app using devise. And should not be used in production at the given time.

# Installation

    go get github.com/skunkworker/go-devise-encryptor
Import

    import "github.com/skunkworker/go-devise-encryptor"

# Usage

If no salt present

```go

package main

import (
  "fmt"
  "github.com/skunkworker/go-devise-encryptor"
)

func main() {
  password := "changeme"
  stretches := 10
  pepper := "a really bad pepper"

  err, hashedPassword := devisecrypto.Digest(password,stretches,pepper)
  if err != nil {
    panic(err)
  }
  fmt.Println("hashedPassword: ",hashedPassword)

  // and to compare with a previously hashed password

  newPassword := "changemeagain"

  err, val := devisecrypto.Compare(newPassword,pepper,hashedPassword)

  if err != nil {
    panic(err)
  }

  if val {
    fmt.Println("Passwords are the same")
  }

}