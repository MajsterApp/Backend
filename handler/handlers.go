package handler

import (
	"net/http"
    "github.com/MajsterApp/Backend/api"

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

func (o *Order) Verification(w http.ResponseWriter, r *http.Request) {
    Verification(w, r)
}

func (o *Order) PasswordChange(w http.ResponseWriter, r *http.Request) {
    PasswordChange(w, r)
}

func (o *Order) FetchCities(w http.ResponseWriter, r *http.Request) {
    api.FetchCities(w, r)
}
//
//     fmt.Println("Read handler")
// }
// func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
//
//     fmt.Println("Delete handler")
// }
// handlers rigt here

