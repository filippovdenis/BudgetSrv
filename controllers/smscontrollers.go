package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func smsPutHandler(w http.ResponseWriter, r *http.Request) {
	println("Method Put")
	vars := mux.Vars(r)

	println("entity = ", entity)
	dec := json.NewDecoder(r.Body)
	switch entity {
	case "category":
		cat := Category{}
		dec.Decode(&cat)
		EntityVar = cat
		println("in category case")
		fmt.Printf("%T: %v\n", EntityVar, EntityVar)
		database.Create(&cat)
	case "smsmessage":
		EntityVar = model.SMSMessage{}
		println("in sms case")
	default:
		EntityVar = nil
	}

	if EntityVar == nil {
		println("nil")
	}
}
