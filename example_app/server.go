package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/endpass/go-endpass"
)

var (
	state        = "123456789"
	clientId     = "<clientId>"
	clientSecret = "<clientSecret>"
	redirectUrl  = "https://example.com/oauth2"
)

func main() {
	client := endpass.NewClient(
		clientId,
		clientSecret,
		[]string{"user:email:read", "documents:image:read"},
		state,
		redirectUrl,
	)

	//-------------------------------------------------------------------------
	// If you need working through proxy, uncomment next block
	//-------------------------------------------------------------------------
	//dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//client.SetDialer(dialer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := client.AuthCodeURL()
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/oauth2", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		state := r.Form.Get("state")
		if !client.IsStateValid(state) {
			http.Error(w, "State invalid", http.StatusBadRequest)
			return
		}
		code := r.Form.Get("code")
		if code == "" {
			http.Error(w, "Code not found", http.StatusBadRequest)
			return
		}
		err = client.Exchange(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/info", http.StatusFound)
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		scopes, err := client.Scopes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := client.User()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = fmt.Fprintf(w, "scopes: %s\n\nemail: %s\n", scopes, user.Email)
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
