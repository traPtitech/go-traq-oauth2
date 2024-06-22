# go-traq-oauth2

[![GoDoc](https://godoc.org/github.com/traPtitech/go-traq-oauth2?status.svg)](https://godoc.org/github.com/traPtitech/go-traq-oauth2)

Package traqoauth2 provides constants for using OAuth2 to access traQ.

## Usage

> [!WARNING]
> OAuth2 Authorization Code Flowに則ったClientを実装する前に、他のより簡単な方法で実装できないか検討してください。
> - ログインユーザーの traQ ID (e.g. `@traP`) を取得したい場合
>   - [traPtitech/NeoShowcase](https://github.com/traPtitech/NeoShowcase)
>   - [traPtitech/caddy-trap-auth](https://github.com/traPtitech/caddy-trap-auth)
>   - [traPtitech/traefik-forward-auth](https://github.com/traPtitech/traefik-forward-auth)
> - ログインユーザーの識別情報を必要としないデータを取得したい場合 (e.g. GET /users, GET /groups/:groupID)
>   - [Bot](https://bot-console.trap.jp/docs/bot)のAccess Tokenを用いる
>   - (実装予定) [OAuth2 Client Credential Flow](https://github.com/traPtitech/traQ/issues/2403) を用いる
> - ログインユーザーの識別情報を利用したい場合 (e.g. GET /users/me, POST /channels/:channelID/messages)
>   - OAuth2 Authorization Code Flow

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
  Scopes:       []string{traqoauth2.ScopeRead, traqoauth2.ScopeWrite}
}
```
