# go-traq-oauth2

[![GoDoc](https://godoc.org/github.com/traPtitech/go-traq-oauth2?status.svg)](https://godoc.org/github.com/traPtitech/go-traq-oauth2)

Package traqoauth2 provides constants for using OAuth2 to access traQ.

## Usage

Full example: [_example/](_example/)

```go
package main

import (
  // ...
  traqoauth2 "github.com/traPtitech/go-traq-oauth2"
)

// Configure at https://bot-console.trap.jp/clients
var  oauth2Config = oauth2.Config{
  ClientID:     os.Getenv("TRAQ_CLIENT_ID"),
  ClientSecret: os.Getenv("TRAQ_CLIENT_SECRET"),
  Endpoint:     traqoauth2.Prod, // or traqoauth2.Staging
  RedirectURL:  os.Getenv("TRAQ_REDIRECT_URL"),
}
```
