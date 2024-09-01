package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServe struct {
	listAddr string
	store    Storage
}

func newAPIServ(listAddr string, store Storage) *APIServe {
	return &APIServe{
		listAddr: listAddr,
		store:    store,
	}
}

func (s *APIServe) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAcc))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAcc))

	log.Println("JSON API RUN on port:", s.listAddr)

	http.ListenAndServe(s.listAddr, router)
}

func (s *APIServe) handleAcc(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAcc(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAcc(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAcc(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServe) handleGetAcc(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	// db.get(id)
	fmt.Println(id)
	//account := NewAcc("Ant", "EZ")
	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServe) handleCreateAcc(w http.ResponseWriter, r *http.Request) error {

	createAccReq := new(CreateAccReq)
	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return err
	}

	account := NewAcc(createAccReq.FName, createAccReq.LName)

	if err := s.store.CreateAcc(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServe) handleDeleteAcc(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServe) handleTrans(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIErr struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadGateway, APIErr{Error: err.Error()})
		}
	}
}
