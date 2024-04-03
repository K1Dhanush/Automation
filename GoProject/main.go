package main

import (
	"net/http"

	"github.com/gorilla/mux"

	db "GoProject/database"
	event "GoProject/eventhandler"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/report/list-all-reports", event.GetAllReports).Methods("GET")
	r.HandleFunc("/report/namespace/create", event.CreateNamespace).Methods("POST")
	http.ListenAndServe("192.168.36.169:8089", r)
}
