package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChecAuth() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Cheking autentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Loggin() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
