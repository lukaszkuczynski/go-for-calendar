package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	//
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	googleOauthConfig, err = google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	gob.Register(&oauth2.Token{})

}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/logout", handleGoogleLogout)
	http.HandleFunc("/callback", handleGoogleCallback)
	http.ListenAndServe(":8080", nil)

}

func handleMain(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if session.Values["token"] != nil {
		val := session.Values["token"]
		var token2 = &oauth2.Token{}
		token2, ok := val.(*oauth2.Token)
		if !ok {
			panic("token was not oauth2.Token type!")
		}
		if token2.Expiry.Before(time.Now()) {
			print("refresh is:")
			print(token2.RefreshToken)
			if token2.RefreshToken == "" {
				http.Redirect(w, r, "/login", http.StatusAccepted)
				return
			}
			token2, err := newTokenFromRefreshToken(token2.RefreshToken, w, r)
			if err != nil {
				panic("Refresh token problem")
			}
			saveToken(r, token2, w)
		}
		eventsText, err := getMyEvents(token2)
		if err == nil {
			fmt.Fprintf(w, eventsText)
			fmt.Fprintf(w, "<a href=\"/logout\">log out </a>")
		}
	} else {

		var htmlIndex = `<html>
<body>
	<a href="/login">Google Log In</a>
</body>
</html>`
		fmt.Fprintf(w, htmlIndex)
	}

}

func newTokenFromRefreshToken(refreshToken string, w http.ResponseWriter, r *http.Request) (*oauth2.Token, error) {
	token := new(oauth2.Token)
	token.RefreshToken = refreshToken
	token.Expiry = time.Now()

	// TokenSource will refresh the token if needed (which is likely in this
	// use case)
	oauthConfig := *googleOauthConfig
	ts := oauthConfig.TokenSource(oauth2.NoContext, token)

	return ts.Token()
}

var (
	oauthStateString = "pseudo-random"
)

var store = sessions.NewCookieStore([]byte("123"))

func handleGoogleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Values["token"] = nil
	var err = sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url := "/"
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleOauthConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar.events.readonly"}
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	// content, err := getMyEvents(r.FormValue("state"), r.FormValue("code"))
	_, err := storeToken(w, r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func storeToken(w http.ResponseWriter, r *http.Request) (bool, error) {
	var state = r.FormValue("state")
	var code = r.FormValue("code")
	if state != oauthStateString {
		return false, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return false, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	saveToken(r, token, w)
	return true, nil
}

func saveToken(r *http.Request, token *oauth2.Token, w http.ResponseWriter) {
	session, _ := store.Get(r, "session.id")
	session.Values["token"] = token
	session.Save(r, w)
}

func getMyEvents(token *oauth2.Token) (string, error) {

	ctx := context.Background()
	var src oauth2.TokenSource = googleOauthConfig.TokenSource(ctx, token)
	srv, err := calendar.NewService(ctx, option.WithTokenSource(src))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	var response string = ""
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
			response += item.Summary
			response += "\n"
		}
	}
	// contents := []byte(response)
	return response, nil

}
