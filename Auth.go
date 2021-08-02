package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// base url
func createBaseUrl(subDomain string) string {
	return fmt.Sprintf("https://%v.zendesk.com/api/v2/", subDomain)

}

// basic Authentication
func basicAuth(user string, pswd string) string {
	auth := user + ":" + pswd
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// get keys from dotenv
func login() (string, string, string) {

	subDomain := os.Getenv("SUBDOMAIN")
	if subDomain == "" {
		log.Fatalln("login failed: to retrieve subdomain from ENV file.")
	}
	email := os.Getenv("EMAIL")
	if email == "" {
		log.Fatalln("login failed: to retrieve email from ENV file.")
	}
	pswd := os.Getenv("PSWD")
	if pswd == "" {
		log.Fatalln("login failed: to retrieve email from ENV file.")
	}
	return subDomain, email, pswd

}

// Prepare url and client with proper auth
func Login(method string, api string) (*http.Response, error) {
	var subDomain, email, pswd string
	subDomain, email, pswd = login()
	client := &http.Client{}
	base := createBaseUrl(subDomain)
	req, err := http.NewRequest(method, base+api, nil)
	if err != nil {
		return nil, fmt.Errorf("http new request error: %v", err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(email, pswd))
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
