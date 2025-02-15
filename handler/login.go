package handler

import (
	"fmt"
	"net/http"
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Login Works")
}
