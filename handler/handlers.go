package handler

import (
	"net/http"
)

type Order struct{}

func (o *Order) Login(w http.ResponseWriter, r *http.Request) {
    LoginFunc(w, r)
}
func (o *Order) Register(w http.ResponseWriter, r *http.Request) {
    RegisterFunc(w, r)
}
func (o *Order) UserData(w http.ResponseWriter, r *http.Request) {
    UserData(w, r)
}
//
//     fmt.Println("Read handler")
// }
// func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
//
//     fmt.Println("Delete handler")
// }
// handlers rigt here

