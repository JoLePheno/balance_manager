package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Server represents the gRPC server
type Server struct {
	db *sql.DB
}

func (s *Server) Initialize() {
	var err error
	fmt.Printf("Connect to mysdq db\n")
	connectionString := "root:root@tcp(localhost:3306)/mysqlapp"

	if s.db, err = sql.Open("mysql", connectionString); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
		log.Fatal(err)
	}
	file, err := os.Open("../../../mysql.sql")
	if err != nil {
		fmt.Printf("Open file error: %s\n", err.Error())
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	requests := strings.Split(string(content), ";\n")
	for _, request := range requests {
		if _, err := s.db.Exec(request); err != nil {
			fmt.Printf("Connection error: %s\n", err.Error())
		}
	}
}

func (s *Server) Crediter(ctx context.Context, in *Transaction) (*Transaction, error) {
	type user struct {
		Name string `json:"Name"`
		ID int `json:"ID"`
		Solde int `json:"Solde"`
	}
	var u user
	query := fmt.Sprintf("SELECT ID, Name, Solde FROM Users WHERE ID=%d", in.AccountId)
	if err := s.db.QueryRow(query).Scan(&u.ID, &u.Name, &u.Solde); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return nil, err
	}
	u.Solde += int(in.Amount)
	query = fmt.Sprintf("UPDATE Users set Solde=%d WHERE ID=%d", u.Solde, in.AccountId)
	s.db.Exec(query)
	return in, nil
}

func (s *Server) NewUser(ctx context.Context, in *Transaction) (*Transaction, error) {
	log.Printf("FROM client: Receive message in debiter:\n\tid: %s\n\taccountId: %s\n\tdescription: %s\n\tnotes: %s\n\tcurrency: %s\n", in.Id, in.AccountId, in.Description, in.Notes, in.Currency)
	//statement := fmt.Sprintf("INSERT INTO users(email, password) VALUES('%s', '%s')")
	//_, err := db.Exec(statement)
	return in, nil
}

func (s *Server) Debiter(ctx context.Context, in *Transaction) (*Transaction, error) {
	type user struct {
		Name string `json:"Name"`
		ID int `json:"ID"`
		Solde int `json:"Solde"`
	}
	var u user
	query := fmt.Sprintf("SELECT ID, Name, Solde FROM Users WHERE ID=%d", in.AccountId)
	if err := s.db.QueryRow(query).Scan(&u.ID, &u.Name, &u.Solde); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return nil, err
	}
	u.Solde -= int(in.Amount)
	query = fmt.Sprintf("UPDATE Users set Solde=%d WHERE ID=%d", u.Solde, in.AccountId)
	s.db.Exec(query)
	return in, nil
}

func (s *Server) GetAmount(ctx context.Context, in *Transaction) (*Transaction, error) {
	type User struct {
		Name string `json:"Name"`
		ID int `json:"ID"`
		Solde int `json:"Solde"`
	}
	var u User
	query := fmt.Sprintf("SELECT ID, Name, Solde FROM Users WHERE ID=%d", in.AccountId)
	if err := s.db.QueryRow(query).Scan(&u.ID, &u.Name, &u.Solde); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return nil, err
	}
	in.Amount = int64(u.Solde)
	return in, nil
}