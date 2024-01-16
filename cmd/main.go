package main

import (
	"log"
	"net/http"
	"os"

	v1 "github.com/DenisOzindzheDev/pet-project-git-login/internal/api/v1"
)

var GithubClientID = os.Getenv("GITHUB_CLIENT_ID")
var GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")

func main() {
	if GithubClientID == "" || GithubClientSecret == "" || len(GithubClientID) == 0 || len(GithubClientSecret) == 0 {
		log.Fatal("GithubClientID and GithubClientSecret are required")
	}

	http.HandleFunc("/", v1.RootHandler)
	http.HandleFunc("/login/", v1.GithubLoginHandler)
	http.HandleFunc("/github/callback/", v1.GithubCallbackHandler)

	addr := "http://localhost:8080"
	log.Printf("Listening on %s", addr)

	log.Panicf("routing", http.ListenAndServe(":8080", nil)) //server
}
