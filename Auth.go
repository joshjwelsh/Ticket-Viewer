package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Auth struct {
	BaseURL   string
	basicAuth string
	Credential
}

type Credential struct {
	subDomain string
	email     string
	pswd      string
}

func NewAuth() Auth {
	cred, err := NewCredential()
	if err != nil {
		log.Fatal(err)
	}

	return Auth{
		BaseURL:    createBaseUrl(cred.subDomain),
		basicAuth:  basicAuth(cred.email, cred.pswd),
		Credential: cred,
	}

}

func NewCredential() (Credential, error) {

	subDomain := os.Getenv("SUBDOMAIN")
	if subDomain == "" {
		return Credential{}, errors.New("login failed: to retrieve subdomain from ENV file.")
	}
	email := os.Getenv("EMAIL")
	if email == "" {
		return Credential{}, errors.New("login failed: to retrieve email from ENV file.")
	}
	pswd := os.Getenv("PSWD")
	if pswd == "" {
		return Credential{}, errors.New("login failed: to retrieve email from ENV file.")
	}

	return Credential{
		subDomain: subDomain,
		email:     email,
		pswd:      pswd,
	}, nil
}

// base url
func createBaseUrl(subDomain string) string {
	return fmt.Sprintf("https://%v.zendesk.com/api/v2/", subDomain)

}

// basic Authentication
func basicAuth(user string, pswd string) string {
	auth := user + ":" + pswd
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Prepare url and client with proper auth
func (a *Auth) Login(method string, api string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, a.BaseURL+api, nil)
	if err != nil {
		return nil, fmt.Errorf("http new request error: %v", err)
	}
	req.Header.Add("Authorization", "Basic "+a.basicAuth)
	req.Header.Add("Content-Type", "application/json")
	if verbose {
		log.Println(req.URL.String())
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client failed to execute request on url %v : %v", req.URL.String(), err)
	}

	return resp, nil

}
