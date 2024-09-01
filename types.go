package main

import (
	"math/rand"
)

type CreateAccReq struct {
	FName string `json:firstName`
	LName string `json:lastName`
}

type Account struct {
	FName   string `json:"firstName"`
	LName   string `json:"lastName"`
	Number  int64  `json:"number"`
	Balance int64  `json:"balance"`
	ID      int    `json:"id"`
}

func NewAcc(firstN, lastN string) *Account {
	return &Account{
		FName:  firstN,
		LName:  lastN,
		Number: int64(rand.Intn(10000)),
	}
}
