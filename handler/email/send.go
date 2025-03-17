package email

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/resend/resend-go/v2"
)

func Send(w http.ResponseWriter, r *http.Request) {
	var rq struct {
		Html    string `json:"emailHtml"`
		Email   string `json:"email"`
		Subject string `json:"subject"`
	}

	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		http.Error(w, "bad Request", http.StatusUnauthorized)
	}
	apiKey := "re_XXShQEzM_CY3F3zEtjUmZdPoTh5nVmtxK"

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <majsterapp@resend.dev>",
		To:      []string{rq.Email},
		Html:    rq.Html,
		Subject: rq.Subject,
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": sent.Id})
}
