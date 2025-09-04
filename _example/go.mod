module example

go 1.23.0

replace github.com/traPtitech/go-traq-oauth2 => ../

require (
	github.com/gorilla/sessions v1.2.2
	github.com/traPtitech/go-traq v0.0.0-20230720010114-3bada4b8a73a
	github.com/traPtitech/go-traq-oauth2 v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.27.0
)

require github.com/gorilla/securecookie v1.1.2 // indirect
