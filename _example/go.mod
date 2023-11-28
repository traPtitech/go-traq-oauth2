module example

go 1.21.4

replace github.com/traPtitech/go-traq-oauth2 => ../

require (
	github.com/gorilla/sessions v1.2.2
	github.com/traPtitech/go-traq v0.0.0-20230720010114-3bada4b8a73a
	github.com/traPtitech/go-traq-oauth2 v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.15.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	golang.org/x/net v0.19.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
