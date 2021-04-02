package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	c1 := http.Cookie{
		Name:     "cookie1",
		Value:    "first-cookie",
		Expires:  expiration,
		Domain:   "localhost.com",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "cookie2",
		Value:    "second-cookie",
		Expires:  expiration,
		Domain:   "localhost.com",
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	st := "Hello. I made useless cookie for you. your local storage boom zz"
	_, err := w.Write([]byte(st))
	if err != nil {
		log.Fatal(err)
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("cookie1")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func askCookie(w http.ResponseWriter, r *http.Request) {
	st := "show me the cookies"
	_, err := w.Write([]byte(st))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	domain := "localhost.com"
	subDomain1 := "sub1." + domain
	homeDomain := "home." + domain

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", setCookie)
	http.HandleFunc(subDomain1+"/", getCookie)
	http.HandleFunc(homeDomain+"/", askCookie)
	server.ListenAndServe()
}
