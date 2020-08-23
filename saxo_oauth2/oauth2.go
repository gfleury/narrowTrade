package saxo_oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
	"golang.org/x/oauth2"
)

var persistedToken string

func GetToken(ctx context.Context, oauth2cfg *oauth2.Config) (*oauth2.Token, error) {
	var token *oauth2.Token

	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	persistedToken = fmt.Sprintf("%s/.it_token", home)

	if _, err := os.Stat(persistedToken); err == nil {
		tokenJSON, err := ioutil.ReadFile(persistedToken)
		if err != nil {
			return nil, err
		}
		t := oauth2.Token{}
		err = json.Unmarshal(tokenJSON, &t)
		if err != nil {
			return nil, err
		}
		token = &t

		if token.Valid() {
			return token, nil
		}
	}

	codeVerifier, _ := cv.CreateCodeVerifier()
	codeChallenge := codeVerifier.CodeChallengeS256()

	urlParam := oauth2.SetAuthURLParam("redirect_uri", "http://localhost:8888/redirect")

	opts := []oauth2.AuthCodeOption{
		urlParam,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	}

	loginURL := oauth2cfg.AuthCodeURL("y90dsygas98dygoidsahf8sa", opts...)

	redirectURL := oauth2cfg.RedirectURL

	// start a web server to listen on a callback URL
	server := &http.Server{Addr: redirectURL}

	code := ""
	// define a handler that will get the authorization code, call the token endpoint, and close the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cbError := r.URL.Query().Get("error")
		if cbError != "" {
			errorDescription := r.URL.Query().Get("error_description")
			io.WriteString(w, fmt.Sprintf("Error: %s\n", errorDescription))
		}

		// get the authorization code
		code = r.URL.Query().Get("code")
		if code == "" {
			log.Println("snap: Url Param 'code' is missing")
			io.WriteString(w, "Error: could not find 'code' URL parameter\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		// return an indication of success to the caller
		io.WriteString(w, `
		<html>
			<body>
				<h1>Login successful!</h1>
				<h2>You can close this window and return to the CLI.</h2>
			</body>
		</html>`)

		// close the HTTP server
		cleanup(server)
	})

	// set up a listener on the redirect port
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		return nil, err
	}

	err = exec.Command("open", loginURL).Start()
	if err != nil {
		return nil, err
	}
	server.Serve(l)

	/* Perform OAuth2 round trip request and obtain a token */
	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	opts = []oauth2.AuthCodeOption{
		urlParam,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier.String()),
	}
	token, err = oauth2cfg.Exchange(ctx, code, opts...)
	if err != nil {
		return nil, err
	}

	err = PersistToken(token)
	return token, err
}

// cleanup closes the HTTP server
func cleanup(server *http.Server) {
	// we run this as a goroutine so that this function falls through and
	// the socket to the browser gets flushed/closed before the server goes away
	go server.Close()
}

func PersistToken(token *oauth2.Token) error {
	// Persisting Token
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(persistedToken, tokenJSON, 0644)
	if err != nil {
		return err
	}

	// Hack to run non-root
	return os.Chown(persistedToken, 501, 20)
}
