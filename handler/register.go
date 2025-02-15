package handler

import (
    "net/http"
    "fmt"
)

func RegisterFunc(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Register Works")

}
