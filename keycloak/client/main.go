package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientId = "app"
	clientSecret = os.Getenv("CLIENT_SECRET")
)


func main(){
	fmt.Println("clientSecret: " + clientSecret)
	ctx := context.Background()
	
	provider, err := oidc.NewProvider(ctx, "http://" + os.Getenv("KEYCLOAK_HOST") +":8080/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID: clientId,
		ClientSecret: clientSecret,
		Endpoint: provider.Endpoint(),
		RedirectURL: "http://localhost:8081/auth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "magica"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request){
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		fmt.Println(r.URL.Query().Get("code"))
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "failed to exchange token", http.StatusBadRequest)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "no id_token", http.StatusBadRequest)
			return
		}

		resp := struct {
				OAuth2Token *oauth2.Token
				RawIDToken string
			}{
				oauth2Token, rawIDToken,
			}
		
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(data)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}