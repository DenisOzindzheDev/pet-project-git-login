package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("state")
	if err != nil {
		http.Error(w, "Bad Request - cookie 'state' not found", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	if r.URL.Query().Get("state") != state.Value {
		http.Error(w, "Bad Request - invalid state", http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	code := r.URL.Query().Get("code")
	requestBodyMap := map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          code,
	}
	requestJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		log.Fatal("Cannot marshal request")
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON)) //serialized req
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "oAuth Server unavalible", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	respbody, _ := io.ReadAll(resp.Body)

	var ghresp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	json.Unmarshal(respbody, &ghresp)

	userInfo := getGithubUserInfo(ghresp.AccessToken)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(userInfo))
}

func getGithubUserInfo(token string) string {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Fatalf("Cannot Get User %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, _ := io.ReadAll(resp.Body)
	return string(respBody)
}
