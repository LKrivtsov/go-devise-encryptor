# go-devise-encryptor
Devise encryptor written in golang. Useful for integrating a go application with a ruby on rails database. This is just a test to see how easy it is to integrate a go app with an existing rails app using devise. And should not be used in production at the given time.
Ruby's bcrypt implementation uses the 2a algorithm, and golang's bcrypt implementation also uses 2a, so they are compatible.

[![Travis status][build-logo]][build-link]

[build-logo]:https://api.travis-ci.org/skunkworker/go-devise-encryptor.svg
[build-link]:https://travis-ci.org/skunkworker/go-devise-encryptor

# Installation

    go get github.com/skunkworker/go-devise-encryptor
Import

    import "github.com/skunkworker/go-devise-encryptor"

# Usage

If you don't have a salt, just pass the blank string ```""```

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

	hashedPassword, err := devisecrypto.Digest(password, stretches, pepper)
	if err != nil {
		panic(err)
	}
	fmt.Println("hashedPassword: ", hashedPassword)

	// and to compare with a previously hashed password

	newPassword := "changeme"

	val := devisecrypto.Compare(newPassword, pepper, hashedPassword)

	if val {
		fmt.Println("Passwords are the same ")
	}

}
```