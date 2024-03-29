package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/RushikeshMarkad16/AzureAD/config"
	"golang.org/x/oauth2"
)

var Client *http.Client

func HandleLandingPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct{}{}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	authURL := getAuthURL()
// 	http.Redirect(w, r, authURL, http.StatusFound)
// }

// func getAuthURL() string {
// 	oauthConfig, err := adal.NewOAuthConfig("https://login.microsoftonline.com/"+config.TenantID, config.ClientID)
// 	if err != nil {
// 		return err.Error()
// 	}
// }

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found in request", http.StatusBadRequest)
		return
	}

	token, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Access Token: ", token.AccessToken)

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "ID token not found", http.StatusBadRequest)
		return
	}
	fmt.Println("idToken : ", idToken)

	Client = config.OIDCconfig.Client(context.Background(), token)
	resp, err := Client.Get("https://graph.microsoft.com/oidc/userinfo")
	if err != nil {
		http.Error(w, "Invalid access token", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Failed to parse UserInfo", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("./template/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Name": userInfo["name"],
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func exchangeCodeForToken(code string) (*oauth2.Token, error) {
	ctx := context.Background()
	fmt.Println("client: ", config.OIDCconfig.ClientID)
	token, err := config.OIDCconfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func handleLogout(w http.ResponseWriter, r *http.Request) {
// 	// Perform any cleanup or logout tasks if necessary
// 	fmt.Fprintf(w, "Logged out successfully!")
// }
