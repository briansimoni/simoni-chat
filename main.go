package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type message struct {
	Data string
}

var chat []string

func main() {

	chat = make([]string, 0)
	fmt.Println("starting")
	http.HandleFunc("/", home)
	http.HandleFunc("/send", sendMessages)
	http.HandleFunc("/getMessages", getMessages)
	http.ListenAndServe(":4000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	home, _ := ioutil.ReadFile("index.html")
	fmt.Fprint(w, string(home))
}

func sendMessages(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m message
	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println(err.Error)
	}
	defer r.Body.Close()
	chat = append(chat, m.Data)

}

func getMessages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(chat)
}
