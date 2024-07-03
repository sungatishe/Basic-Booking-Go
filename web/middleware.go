package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)
	csrfToken.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProd,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfToken
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
