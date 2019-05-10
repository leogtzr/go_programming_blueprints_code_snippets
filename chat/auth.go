package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

type authHandler struct {
	next http.Handler
}

var (
	providers = map[string]struct{}{
		"google":   {},
		"facebook": {},
		"github":   {},
	}
	actions = map[string]struct{}{
		"login":    {},
		"callback": {},
		"github":   {},
	}
)

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// MustAuth ...
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

/*
/auth/login/google
/auth/login/facebook
/auth/callback/google
/auth/callback/facebook
*/
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")

	if len(segs) < 4 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "actions and providers not specified :(")
		return
	}

	action := segs[2]
	provider := segs[3]

	if _, ok := providers[strings.ToLower(provider)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Provider %s not supported", provider)
		return
	}

	if _, ok := providers[strings.ToLower(provider)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Action %s not supported", action)
		return
	}

	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s: %s", provider, err), http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err), http.StatusInternalServerError)
			return
		}
		authCookieValue := objx.New(map[string]interface{}{
			"name": user.Name(),
			"avatar_url": user.AvatarURL(),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/"})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}


él
