package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountVerifyService struct {
	repo repository.Accounts
}

func NewAccountVerifyService(repo repository.Accounts) *AccountVerifyService {
	return &AccountVerifyService{repo: repo}
}

func (s *AccountVerifyService) LoginApi(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetData(ctx, accountID)
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

	contain, err := CheckAuth(resp)
	if err != nil {
		return err
	}
	if contain {
		return fmt.Errorf("Many request")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Login api: %s", resp.Status)
	}

	var getData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&getData)
	if err != nil {
		return err
	}
	randomHash := getData["random_hash"].(string)
	if err := s.repo.AddRandomHash(ctx, accountID, randomHash); err != nil {
		return err
	}

	return nil
}

func (s *AccountVerifyService) ParsingApi(ctx context.Context, accountLogin domain.AccountLogin) error {
	account, err := s.repo.GetData(ctx, accountLogin.ID)
	if err != nil {
		return err
	}

	data := url.Values{
		"phone":       {account.Phone},
		"random_hash": {account.Random_hash},
		"password":    {accountLogin.Password},
	}

	// Authenticated on website
	client := &http.Client{}
	req, err := http.NewRequest("POST", link_authorized, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	randomIndex := rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	cookie := resp.Cookies()
	defer resp.Body.Close()

	contain, err := CheckAuth(resp)
	if err != nil {
		return err
	}
	if contain {
		return fmt.Errorf("Many request")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Login api: %s", resp.Status)
	}

	// Parsing hash from input
	req, err = http.NewRequest("GET", link_apps, nil)
	if err != nil {
		return err
	}

	for i := range cookie {
		req.AddCookie(cookie[i])
	}
	randomIndex = rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	hash_input := doc.Find(".app_edit_page").Find("input")
	hash, _ := hash_input.Attr("value")
	fmt.Println(hash)

	// Parsing api
	req, err = http.NewRequest("GET", link_apps, nil)
	if err != nil {
		return err
	}

	for i := range cookie {
		req.AddCookie(cookie[i])
	}
	randomIndex = rand.Intn(len(userAgents))
	req.Header.Set("User-Agent", userAgents[randomIndex])

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	api := doc.Find(".app_edit_page").Find(".input-xlarge").Text()
	fmt.Println(api)

	return nil
}

func CheckAuth(resp *http.Response) (bool, error) {
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	contain := strings.Contains(string(response), error_many_request)
	if contain {
		return false, nil
	}
	return true, nil
}

func CreateApp(cookies http.Cookie, hash string) error {
	return nil
}

func (s *AccountVerifyService) GetCodeSession(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetData(ctx, accountID)
	if err != nil {
		return err
	}
	fmt.Println(account)

	return nil
}

func (s *AccountVerifyService) CreateSession(ctx context.Context, accountLogin domain.AccountLogin) error {
	account, err := s.repo.GetData(ctx, accountLogin.ID)
	if err != nil {
		return err
	}
	fmt.Println(account)

	return nil
}
