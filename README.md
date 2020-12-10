# GoZeroBounce
Go implementation for [ZeroBounce Email Validation API](https://www.zerobounce.net/docs/email-validation-api-quickstart/) 

THIS PROJECT IS STILL IN DEVELOPMENT!

## Installation and Usage
```sh
go get github.com/antsanchez/gozerobounce
```

You can use it like this:
```go
package main

import (
    "fmt"
    "os"

    "github.com/antsanchez/gozerobounce"
)

func main() {

    gozerobounce.APIKey = "... Your API KEY ..." 

    // For Querying a single E-Mail and IP
    // IP can also be an empty string
    response := gozerobounce.Validate("email@example.com", "123.123.123.123")

    // You can check in the response the different attributes, 
    // or use methods like those:

    if response.IsDisposable() {
        fmt.Println("This E-mail is disposable")
    }

    if response.IsValid() {
        fmt.Println("This E-mail is valid")
    }

    if response.IsFreeEmail() {
        fmt.Println("This E-mail comes from a free provider")
    }

    // You can also check your credits 
    credits := gozerobounce.GetCredits()
    fmt.Println("Credits left:", credits.Credits)
}
```

## Already implemented
- Validate single email
- Get gredits
