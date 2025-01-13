package main

import (
	"encoding/json"
	"net/http"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	// var service KeycloakService
	// res := service.login(new(KLoginPayload))
	// fmt.Println(res)

	payload := new(LoginPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invilid payload"))
		return
	}

	kpayload := &KLoginPayload{
		clientId:     "rest-golang-auth",
		username:     payload.Username,
		password:     payload.Password,
		grantType:    "password",
		clientSecret: "o2AgY8QqPicQI5iJWUdBpt76ufGnrgsY",
	}

	kres, err := s.client.login(kpayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res := &LoginRes{
		AccessToken: kres.AccessToken,
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
