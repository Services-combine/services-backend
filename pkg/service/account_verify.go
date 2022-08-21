package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/logging"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	link_get_password     = "https://my.telegram.org/auth/send_password"
	link_authorized       = "https://my.telegram.org/auth/login"
	link_apps             = "https://my.telegram.org/apps"
	link_create_app       = "https://my.telegram.org/apps/create"
	error_many_request    = "Sorry, too many tries. Please try again later."
	error_invalid_code    = "Invalid confirmation code!"
	symbols               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers               = "1234567890"
	path_python           = "/usr/bin/python3"
	script_send_code      = "send_code.py"
	script_verify_account = "verify_account.py"
)

type AccountVerifyService struct {
	repo repository.Accounts
}

func NewAccountVerifyService(repo repository.Accounts) *AccountVerifyService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &AccountVerifyService{repo: repo}
}

func (s *AccountVerifyService) LoginApi(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetById(ctx, accountID)
	if err != nil {
		return err
	}

	data := url.Values{
		"phone": {account.Phone},
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", link_get_password, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch string(response) {
	case error_many_request:
		return fmt.Errorf("Many request")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Authentication: %s", resp.Status)
	}

	var getData map[string]interface{}
	json.Unmarshal(response, &getData)
	randomHash := getData["random_hash"].(string)
	if err := s.repo.AddRandomHash(ctx, accountID, randomHash); err != nil {
		return err
	}

	return nil
}

func (s *AccountVerifyService) ParsingApi(ctx context.Context, accountLogin domain.AccountLogin) error {
	var accountApi domain.AccountApi
	accountApi.ID = accountLogin.ID

	account, err := s.repo.GetById(ctx, accountLogin.ID)
	if err != nil {
		return err
	}
	client := &http.Client{}

	cookie, err := AuthenticationWebsite(account, accountLogin.Password, client)
	if err != nil {
		return err
	}

	hash, err := ParsingHashInput(cookie, client)
	if err != nil {
		return err
	}

	err = CreateApp(cookie, client, hash)
	if err != nil {
		return err
	}

	api_id, api_hash, err := ParsingApiApp(cookie, client)
	if err != nil {
		return err
	}

	accountApi.ApiId = api_id
	accountApi.ApiHash = api_hash
	if err := s.repo.AddApi(ctx, accountApi); err != nil {
		return err
	}

	return nil
}

func AuthenticationWebsite(account domain.Account, password string, client *http.Client) ([]*http.Cookie, error) {
	data := url.Values{
		"phone":       {account.Phone},
		"random_hash": {account.Random_hash},
		"password":    {password},
	}
	var cookie []*http.Cookie

	req, err := http.NewRequest("POST", link_authorized, strings.NewReader(data.Encode()))
	if err != nil {
		return cookie, err
	}

	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return cookie, err
	}
	cookie = resp.Cookies()
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cookie, err
	}

	switch string(response) {
	case error_many_request:
		return cookie, fmt.Errorf("Many request")
	case error_invalid_code:
		return cookie, fmt.Errorf("Invalid code")
	}

	if resp.StatusCode != http.StatusOK {
		return cookie, fmt.Errorf("Authentication: %s", resp.Status)
	}

	return cookie, nil
}

func ParsingHashInput(cookie []*http.Cookie, client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", link_apps, nil)
	if err != nil {
		return "", err
	}

	for i := range cookie {
		req.AddCookie(cookie[i])
	}
	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	hash_input := doc.Find(".app_edit_page").Find("input")
	hash, _ := hash_input.Attr("value")
	return hash, nil
}

func GenerateSymbols(n int) string {
	sequence := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range sequence {
		sequence[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(sequence)
}

func CreateApp(cookies []*http.Cookie, client *http.Client, hash string) error {
	app_title := GenerateSymbols(9)
	app_shortname := GenerateSymbols(7)

	app_title += fmt.Sprint(rand.Intn(100))
	app_shortname += fmt.Sprint(rand.Intn(10))

	data := url.Values{
		"hash":          {hash},
		"app_title":     {app_title},
		"app_shortname": {app_shortname},
		"app_url":       {""},
		"app_platform":  {"Android"},
		"app_desc":      {""},
	}

	req, err := http.NewRequest("POST", link_create_app, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func ParsingApiApp(cookie []*http.Cookie, client *http.Client) (int, string, error) {
	req, err := http.NewRequest("GET", link_apps, nil)
	if err != nil {
		return 0, "", err
	}

	for i := range cookie {
		req.AddCookie(cookie[i])
	}
	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, "", err
	}

	index := 1
	var api_id_s string
	var api_hash string
	doc.Find(".app_edit_page").Find(".input-xlarge").Each(func(i int, selection *goquery.Selection) {
		if index == 1 {
			api_id_s = selection.Text()
		} else if index == 2 {
			api_hash = selection.Text()
		}
		index++
	})

	api_id, err := strconv.Atoi(api_id_s)
	if err != nil {
		return 0, "", err
	}

	return api_id, api_hash, nil
}

func (s *AccountVerifyService) GetCodeSession(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetById(ctx, accountID)
	if err != nil {
		return err
	}

	script := os.Getenv("FOLDER_PYTHON_SCRIPTS_VERIFY") + script_send_code
	args_phone := fmt.Sprintf("-P %s", account.Phone)
	args_hash := fmt.Sprintf("-H %s", account.Api_hash)
	args_id := fmt.Sprintf("-I %d", account.Api_id)

	phone_code_hash, err := exec.Command(path_python, script, args_phone, args_hash, args_id).Output()
	if err != nil {
		return err
	}

	if string(phone_code_hash) == "ERROR" {
		return fmt.Errorf("Ошибка при получении кода")
	}

	if err := s.repo.AddPhoneHash(ctx, accountID, string(phone_code_hash)); err != nil {
		return err
	}

	return nil
}

func (s *AccountVerifyService) CreateSession(ctx context.Context, accountLogin domain.AccountLogin) error {
	account, err := s.repo.GetById(ctx, accountLogin.ID)
	if err != nil {
		return err
	}

	script := os.Getenv("FOLDER_PYTHON_SCRIPTS_VERIFY") + script_verify_account
	args_phone := fmt.Sprintf("-P %s", account.Phone)
	args_hash := fmt.Sprintf("-H %s", account.Api_hash)
	args_id := fmt.Sprintf("-I %d", account.Api_id)
	args_code := fmt.Sprintf("-C %s", accountLogin.Password)
	args_codeHash := fmt.Sprintf("-G %s", account.Phone_code_hash)

	result, err := exec.Command(path_python, script, args_phone, args_hash, args_id, args_code, args_codeHash).Output()
	if err != nil {
		return err
	}

	if string(result) == "ERROR" {
		return fmt.Errorf("Ошибка при создании .session файла")
	}

	if err := s.repo.ChangeVerify(ctx, accountLogin.ID); err != nil {
		return err
	}

	return nil
}
