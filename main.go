package main

import (
	"fmt"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/filippovdenis/BudgetSrv/controllers"
	"github.com/filippovdenis/BudgetSrv/models"
)

var database *gorm.DB

func getHandler(w http.ResponseWriter, r *http.Request) {
	println("Method Get")
	vars := mux.Vars(r)
	id := vars["id"]
	entity := vars["entity"]
	println("entity = ", entity)
	println("id = ", id)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	println("Method Post")
	vars := mux.Vars(r)
	id := vars["id"]
	entity := vars["entity"]
	println("entity = ", entity)
	println("id = ", id)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	println("Method Put")
	vars := mux.Vars(r)
	var EntityVar interface{}

	entity := vars["entity"]
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

func main() {
	fmt.Println("hello")

	connStr := "user=postgres password=tamplier dbname=TestDB sslmode=disable"

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	defer db.Close()
	smsPut

	router := mux.NewRouter()
	router.HandleFunc("/smsmessage", smsPutHandler).Methods("PUT")
	router.HandleFunc("/{entity}/{id:[0-9]+}", getHandler).Methods("GET")
	router.HandleFunc("/{entity}/{id:[0-9]+}", postHandler).Methods("POST")
	router.HandleFunc("/{entity}", putHandler).Methods("PUT")
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}
