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
	email := os.Getenv("EMAIL")
	pswd := os.Getenv("PSWD")
	// Error on missing credentials
	if subDomain == "" || email == "" || pswd == "" {
		log.Panicln("Subdomain, email, or password missing from .env file.")
	}
	return subDomain, email, pswd
}

func Login(method string, api string) *http.Response {
	subDomain, email, pswd := login()
	client := &http.Client{}
	base := createBaseUrl(subDomain)
	req, err := http.NewRequest(method, base+api, nil)
	if err != nil {
		log.Fatalln("Error - creating http new request - ", err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(email, pswd))
	req.Header.Add("Content-Type", "application/json")
	log.Println(req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error - creating http new response - ", err)
	}
	return resp

}
