package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type MessagesResponse struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type Server struct {
	Messages []*Message
}

func (s *Server) AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m Message
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}
	s.Messages = append(s.Messages, &m)
	w.Write([]byte("OK"))
}

func (s *Server) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.Messages)
}

func main() {
	messages := make([]*Message, 0)
	s := &Server{
		Messages: messages,
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", s.AddMessageHandler).Methods("POST")
	r.HandleFunc("/", s.GetMessagesHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}
