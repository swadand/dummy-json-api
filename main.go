package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	Model "github.com/swadand/dummy-json-api/model"
)

func main() {
	userList := Model.UserList()

	fmt.Println("Server started!")
	http.Handle("/", http.FileServer(http.Dir("./views")))
	http.HandleFunc("/json/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userList)

	})

	http.ListenAndServe(":8000", nil)
}
