package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

type UserClient struct {
	ID   string
	Name string
}

type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}

type UserServer struct {
	ID string
}

type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

func main() {
	var net bytes.Buffer //dummy network

	clientTx := &TxClient{
		ID: "txClient",
		User: &UserClient{
			ID:   "ClientID",
			Name: "Client",
		},
		AccountFrom: "Bob",
		AccountTo:   "Pietje",
		Amount:      9.99,
	}

	enc := gob.NewEncoder(&net) //encode the data and the target for the data is the dummy network
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("Error encoding: ", err)
	}
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("Server error:", err)
	}

	fmt.Printf("%#v\n", serverTx)
}

func sendToServer(net io.Reader) (*TxServer, error) {
	tx := &TxServer{}
	dec := gob.NewDecoder(net)
	//err :=
	return tx, dec.Decode(tx)
}
