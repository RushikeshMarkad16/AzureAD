package config

import (
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
	"golang.org/x/oauth2"
)

var (
	OIDCconfig   *oauth2.Config
	ClientID     string
	ClientSecret string
	TenantID     string
	RedirectURL  string
)

func loadEnvVars() {
	ClientID = os.Getenv("ClientID")
	ClientSecret = os.Getenv("ClientSecret")
	TenantID = os.Getenv("TenantID")
	RedirectURL = os.Getenv("RedirectURL")
}

func AzureOauthConfig() {
	OIDCconfig = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/9a09bc79-207d-410b-9c33-ef906c4c06b0/oauth2/v2.0/authorize",
			TokenURL: "https://login.microsoftonline.com/9a09bc79-207d-410b-9c33-ef906c4c06b0/oauth2/v2.0/token",
		},
		Scopes: []string{"openid", "profile", "email"},
	}
}

// func AzureSAMLConfig() {}

// Load ...
func Load() {
	err := gotenv.Load()
	if err != nil {
		fmt.Println("error : ", err)
		log.Fatal("Error loading .env file")
	}

	loadEnvVars()
	AzureOauthConfig()
	// AzureSAMLConfig()
}
