package main

import (
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/traPtitech/go-traq"
	traqoauth2 "github.com/traPtitech/go-traq-oauth2"
	"golang.org/x/oauth2"
)

const (
	sessionName = "session_name"
	stateLength = 16
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     os.Getenv("TRAQ_CLIENT_ID"),
		ClientSecret: os.Getenv("TRAQ_CLIENT_SECRET"),
		Endpoint:     traqoauth2.Prod,
		RedirectURL:  os.Getenv("TRAQ_REDIRECT_URL"), // e.g. http://localhost:8080/oauth2/callback,
		Scopes:       []string{traqoauth2.ScopeRead, traqoauth2.ScopeWrite},
	}

	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
)

func main() {
	// Register oauth2.Token to gob for session
	gob.Register(&oauth2.Token{})

	server := http.NewServeMux()

	server.HandleFunc("/oauth2/authorize", authorizeHandler)
	server.HandleFunc("/oauth2/callback", callbackHandler)
	server.HandleFunc("/me", getMeHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		internalHTTPError(w, err, "failed to get session")
		return
	}

	codeVerifier := oauth2.GenerateVerifier()
	session.Values["code_verifier"] = codeVerifier

	state := make([]byte, stateLength)
	_, _ = rand.Read(state)
	session.Values["state"] = string(state)

	if err := session.Save(r, w); err != nil {
		internalHTTPError(w, err, "failed to save session")
		return
	}

	// this client forces to use PKCE
	// code_challenge_method = S256 is set by S256ChallengeOption
	authCodeURL := oauth2Config.AuthCodeURL(
		string(state),
		oauth2.S256ChallengeOption(codeVerifier),
	)

	http.Redirect(w, r, authCodeURL, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// query parameters
	var (
		code  = r.URL.Query().Get("code")
		state = r.URL.Query().Get("state")
	)

	if code == "" {
		http.Error(w, "code is empty", http.StatusBadRequest)
		return
	}

	session, err := store.Get(r, sessionName)
	if err != nil {
		internalHTTPError(w, err, "failed to get session")
		return
	}

	codeVerifier, ok := session.Values["code_verifier"]
	if !ok {
		http.Error(w, "invalid session", http.StatusBadRequest)
		return
	}

	if storedState, ok := session.Values["state"]; !ok || storedState.(string) != state {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	token, err := oauth2Config.Exchange(
		r.Context(),
		code,
		oauth2.VerifierOption(codeVerifier.(string)),
	)
	if err != nil {
		internalHTTPError(w, err, "failed to exchange token")
		return
	}

	session.Values["token"] = token
	if err := session.Save(r, w); err != nil {
		internalHTTPError(w, err, "failed to save session")
		return
	}

	_, _ = w.Write([]byte("success! you can close this window."))
}

func getMeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		internalHTTPError(w, err, "failed to get session")
		return
	}

	token, ok := session.Values["token"].(*oauth2.Token)
	if !ok {
		http.Error(w, "invalid session", http.StatusBadRequest)
		return
	}

	traqConfig := traq.NewConfiguration()
	traqConfig.HTTPClient = oauth2Config.Client(r.Context(), token)
	client := traq.NewAPIClient(traqConfig)

	user, res, err := client.MeApi.GetMe(r.Context()).Execute()
	if err != nil {
		internalHTTPError(w, err, "failed to get me")
		return
	}
	if res.StatusCode != http.StatusOK {
		http.Error(w, "failed to get me", res.StatusCode)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Hello, %s!", user.Name)))
}

func internalHTTPError(w http.ResponseWriter, err error, msg string) {
	log.Printf("%s: %v", msg, err)
	http.Error(w, msg, http.StatusInternalServerError)
}
