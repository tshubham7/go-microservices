package main

import (
	"log"
	"net"
	"os"

	"github.com/tshubham7/go-microservices/invoice/db"
	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/invoice/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var l = log.New(os.Stdout, "invoice-service ", log.LstdFlags)

func main() {

	wdb, err := db.GetDatabase()
	if err != nil {
		l.Fatal(err)
	}

	db.RunMigrations(wdb)
	l.Println("successfully migrated models")

	gs := grpc.NewServer()
	is := server.NewInvoice(wdb)

	protos.RegisterInvoiceServer(gs, is)

	reflection.Register(gs)

	nt, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	l.Println("listening to port :9002")
	gs.Serve(nt)
}
