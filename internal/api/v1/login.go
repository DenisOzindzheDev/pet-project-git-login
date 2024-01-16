package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DenisOzindzheDev/pet-project-git-login/pkg/helpers"
)

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := helpers.RandString(16)
	if err != nil {
		log.Panic(err)
	}
	c := &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&state=%s", os.Getenv("GITHUB_CLIENT_ID"), state)
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}
