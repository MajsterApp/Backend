package handler

import (
	"net/http"
    "github.com/MajsterApp/Backend/handler/email"

)

type Order struct{}

func (o *Order) LoginHandler(w http.ResponseWriter, r *http.Request) {
    LoginFunc(w, r)
}
func (o *Order) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    RegisterFunc(w, r)
}
func (o *Order) UserDataHandler(w http.ResponseWriter, r *http.Request) {
    UserData(w, r)
}

func (o *Order) VerificationHandler(w http.ResponseWriter, r *http.Request) {
    Verification(w, r)
}

func (o *Order) PasswordChangeHandler(w http.ResponseWriter, r *http.Request) {
    PasswordChange(w, r)
}

func (o *Order) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
    CreateToken(w, r)
}
func (o *Order) GetCitiesHandler(w http.ResponseWriter, r *http.Request) {
    GetCities(w, r)
}
func (o *Order) SendEmail(w http.ResponseWriter, r *http.Request) {
    email.Send(w, r)
}
// func (o *Order) FetchCities(w http.ResponseWriter, r *http.Request) {
//     api.FetchCities(w, r)
// }
//
//     fmt.Println("Read handler")
// }
// func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
//
//     fmt.Println("Delete handler")
// }
// handlers rigt here

