package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAcc(*Account) error
	DeleteAcc(int) error
	UpdateAcc(*Account) error
	getAccById(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {
	conn := "user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{db: db}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccTable()
}

func (s *PostgressStore) createAccTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    number SERIAL UNIQUE,
    balance NUMERIC(15, 2) DEFAULT 0
);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAcc(acc *Account) error {
	query := `
		insert into account 
		values(
			$1,$2,$3,$4,$5
		)
	`
	res, err := s.db.Query(query,
		acc.Balance,
		acc.FName,
		acc.LName,
		acc.ID,
		acc.Number,
	)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func (s *PostgressStore) UpdateAcc(*Account) error {
	return nil
}
func (s *PostgressStore) DeleteAcc(id int) error {
	return nil
}
func (s *PostgressStore) getAccById(id int) (*Account, error) {
	return nil, nil
}
