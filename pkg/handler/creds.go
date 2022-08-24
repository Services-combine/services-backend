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
	"time"

	"github.com/korpgoodness/service.git/internal/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func (h *Handler) GetClient(ctx context.Context, appTokenPath, userTokenFile string) (*http.Client, error) {
	fileBytes, err := ioutil.ReadFile(appTokenPath)
	if err != nil {
		return nil, err
	}

	config, err := h.GenerateConfig(fileBytes)
	if err != nil {
		return nil, err
	}

	token, err := h.ReadUserToken(userTokenFile)
	if err != nil || !token.Valid() || time.Until(token.Expiry).Hours() <= 1 {
		token, err = h.GetUserTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		token.Expiry = token.Expiry.AddDate(0, 3, 0)

		var mapToken map[string]interface{}
		bytesToken, _ := json.Marshal(token)
		json.Unmarshal(bytesToken, &mapToken)
		mapToken["client_id"] = config.ClientID
		mapToken["client_secret"] = config.ClientSecret

		err = h.SaveUserToken(userTokenFile, mapToken)
		if err != nil {
			return nil, err
		}
	}

	return config.Client(ctx, token), nil
}

func (h *Handler) ReadAppToken(appTokenPath string) ([]byte, error) {
	b, err := ioutil.ReadFile(appTokenPath)
	return b, err
}

func (h *Handler) GenerateConfig(fileBytes []byte) (*oauth2.Config, error) {
	config, err := google.ConfigFromJSON(fileBytes, youtube.YoutubeForceSslScope, youtube.YoutubeUploadScope)
	if err != nil {
		return nil, err
	}
	config.RedirectURL = "http://" + os.Getenv("URL_LISTEN_OAUTH_CODE")

	return config, nil
}

func (h *Handler) ReadUserToken(userTokenFile string) (*oauth2.Token, error) {
	f, err := os.Open(userTokenFile)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func (h *Handler) GetUserTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	codeCh, err := h.StartWebServer()
	if err != nil {
		return nil, err
	}

	err = h.OpenURL(authURL)
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

func (h *Handler) StartWebServer() (codeCh chan string, err error) {
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

func (h *Handler) OpenURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = domain.ErrUnableOpenUrl
	}
	return err
}

func (h *Handler) SaveUserToken(userTokenFile string, token map[string]interface{}) error {
	f, err := os.OpenFile(userTokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}
