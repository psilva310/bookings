package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Test middleware. It writes to console every time you load the page
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurt adds CSRF protection to all POST requets
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.SslEnabled,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads ands saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
