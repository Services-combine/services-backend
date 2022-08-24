package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func getClient(ctx context.Context, appTokenPath, userTokenFile string) (*http.Client, error) {
	fileBytes, err := ioutil.ReadFile(appTokenPath)
	if err != nil {
		return nil, err
	}

	config, err := generateConfig(fileBytes)
	if err != nil {
		return nil, err
	}

	token, err := readUserToken(userTokenFile)
	if err != nil || !token.Valid() {
		token, err = getUserTokenFromWeb(config)
		if err != nil {
			return nil, err
		}

		token.Expiry = token.Expiry.AddDate(0, 1, 0)
		err = saveUserToken(userTokenFile, token)
		if err != nil {
			return nil, err
		}
	}
	return config.Client(ctx, token), nil
}

func readAppToken(appTokenPath string) ([]byte, error) {
	b, err := ioutil.ReadFile(appTokenPath)
	return b, err
}

func generateConfig(fileBytes []byte) (*oauth2.Config, error) {
	config, err := google.ConfigFromJSON(fileBytes, youtube.YoutubeForceSslScope, youtube.YoutubeUploadScope)
	if err != nil {
		return nil, err
	}
	config.RedirectURL = "http://" + os.Getenv("URL_LISTEN_OAUTH_CODE")

	return config, nil
}

func readUserToken(userTokenFile string) (*oauth2.Token, error) {
	f, err := os.Open(userTokenFile)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func getUserTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	codeCh, err := startWebServer()
	if err != nil {
		return nil, err
	}

	err = openURL(authURL)
	if err != nil {
		return nil, err
	}

	code := <-codeCh
	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func startWebServer() (codeCh chan string, err error) {
	listener, err := net.Listen("tcp", os.Getenv("URL_LISTEN_OAUTH_CODE"))
	if err != nil {
		return nil, err
	}
	codeCh = make(chan string)

	go http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		codeCh <- code
		listener.Close()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Received code: %v\r\nYou can now safely close this browser window.", code)
	}))

	return codeCh, nil
}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}

func saveUserToken(userTokenFile string, token *oauth2.Token) error {
	f, err := os.OpenFile(userTokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}
