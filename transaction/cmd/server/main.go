package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"balance_manager/balance/api"
)

type _Transaction struct {
	id int64
	accountId int64
	description string
	currency string
	notes string
	param string
	amount int64
}

func main() {
	var cfg _Transaction
	flag.Int64Var(&cfg.id, "id", 0, "id of the transaction")
	flag.Int64Var(&cfg.accountId, "accountId", 0, "Account Id")
	flag.StringVar(&cfg.description, "description", "", "Transaction description")
	flag.StringVar(&cfg.currency, "currency", "", "Transaction currency")
	flag.StringVar(&cfg.notes, "notes", "", "Transaction notes")
	flag.Int64Var(&cfg.amount, "amount", 0, "amount of the transaction")
	flag.StringVar(&cfg.param, "param", "", "debiter / crediter / getAmount")
	flag.Parse()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)
	var response *api.Transaction
	req := api.Transaction{Id: cfg.id, AccountId: cfg.accountId, Description: cfg.description, Currency: cfg.currency, Notes: cfg.notes, Amount: cfg.amount}
	if cfg.param == "getAmount" {
		if response, err = c.GetAmount(context.Background(), &req); err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Printf("Solde: %d\n", response.Amount)
	}
	if cfg.param == "" {
		return
	}
	if cfg.param == "crediter" {
		if response, err = c.Crediter(context.Background(), &req); err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	}
	if cfg.param == "debiter" {
		if response, err = c.Debiter(context.Background(), &req); err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	}
	fmt.Printf("operation :%s succeed", response.Description)
}
