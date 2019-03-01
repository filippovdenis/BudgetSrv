package controllers

import (
	"encoding/json"
	"github.com/filippovdenis/BudgetSrv/models"
	"net/http"
)

func smsPutHandler(w http.ResponseWriter, r *http.Request) {
	println("Method Put")

	dec := json.NewDecoder(r.Body)
	sms := models.SMSMessage{}
	dec.Decode(&sms)

}
