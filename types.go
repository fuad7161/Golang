package main

import "net/http"

type KeycloakService interface {
	login(*KLoginPayload) (*KloginRes, error)
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type KLoginPayload struct {
	clientId     string
	username     string
	password     string
	grantType    string
	clientSecret string
}

type Client struct {
	httpClient *http.Client
}

type KloginRes struct {
	AccessToken string `json:"access_token"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}
