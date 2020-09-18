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

var lg = log.New(os.Stdout, "invoice-service ", log.LstdFlags)

func main() {

	wdb, err := db.GetDatabase()
	if err != nil {
		lg.Fatal(err)
	}

	db.RunMigrations(wdb)
	lg.Println("successfully migrated models")

	gs := grpc.NewServer()
	is := server.NewInvoice(wdb, lg)

	protos.RegisterInvoiceServer(gs, is)

	reflection.Register(gs)

	nt, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	lg.Println("listening to port :9002")
	gs.Serve(nt)
}
