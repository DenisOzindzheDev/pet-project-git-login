package v1

import "net/http"

const rootHTML = `
<h1>Test</h1>
<p>Using raw HTTP OAuth 2.0</p>
<p>You can log into this app with your GitHub credentials:</p>
<p><a href="/login/">Log in with GitHub</a></p>`

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(rootHTML))
}
