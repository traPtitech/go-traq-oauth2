package traqoauth2

import (
	"net/url"

	"golang.org/x/oauth2"
)

const (
	authorizePath = "/oauth2/authorize"
	tokenPath     = "/oauth2/token"
)

var (
	Prod, _    = New("https://q.trap.jp/api/v3")
	Staging, _ = New("https://q-dev.trapti.tech/api/v3")
)

// New returns an Endpoint for the given API base URL.
func New(apiBaseURL string) (oauth2.Endpoint, error) {
	parsedURL, err := url.Parse(apiBaseURL)
	if err != nil {
		return oauth2.Endpoint{}, err
	}

	return oauth2.Endpoint{
		AuthURL:  parsedURL.JoinPath(authorizePath).String(),
		TokenURL: parsedURL.JoinPath(tokenPath).String(),
	}, nil
}
