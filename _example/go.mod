module example

go 1.25.0

replace github.com/traPtitech/go-traq-oauth2 => ../

require (
	github.com/gorilla/sessions v1.4.0
	github.com/traPtitech/go-traq v0.0.0-20260603150909-e1835f95a651
	github.com/traPtitech/go-traq-oauth2 v1.0.0
	golang.org/x/oauth2 v0.36.0
)

require (
	github.com/gorilla/securecookie v1.1.2 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
)
