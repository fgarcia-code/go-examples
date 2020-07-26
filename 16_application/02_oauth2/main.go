package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"context"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "client-id",
		ClientSecret: "14d3e2g4-s43f-1t25-85d7-1h53ca27e00g",
		RedirectURL:  "http://localhost:8090/auth/callback",
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL: "http://localhost:8080/auth/realms/TEST/protocol/openid-connect/auth",
			TokenURL: "http://localhost:8080/auth/realms/TEST/protocol/openid-connect/token",
		},
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	fmt.Printf("\nInsert the auth code:")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AccessToken")
	fmt.Println(tok.AccessToken)
	fmt.Println("RefreshToken")
	fmt.Println(tok.RefreshToken)
	client := conf.Client(ctx, tok)
	client.Get("...")
}
