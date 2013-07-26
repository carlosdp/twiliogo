# twilio-go
The unofficial Go helper library for [Twilio](http://twilio.com).

# Installation

``` bash
go get github.com/carlosdp/twilio-go
```

# Usage

## Send a Text

``` go
package main

import (
  "fmt"
  "github.com/carlosdp/twilio-go
)

func main() {
  client := twilio.NewClient("<ACCOUNT_SID", "<AUTH_TOKEN>")

  message, err := twilio.SendMessage(client, "3334445555", "2223334444", "Hello World!")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Message sent!")
  }
}
```

## Make a Call

``` go
package main

import (
  "fmt"
  "github.com/carlosdp/twilio-go"
)

func main() {
  client := twilio.NewClient("<ACCOUNT_SID>", "<AUTH_TOKEN>")

  call, err := twilio.MakeCall(client, "8883332222", "3334443333")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Call Queued!")
  }
}
```
