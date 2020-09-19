package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/go-microservices/invoice/db"
	"github.com/tshubham7/go-microservices/invoice/handler"
	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/invoice/repository"
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

	lg.Println("grpc listening to port :5000")
	nt, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	go gs.Serve(nt)

	gn := gin.Default()

	ir := repository.NewInvoiceRepo(wdb, lg)

	invoice(gn, ir)

	gn.Run(":9002")

}

// invoice routes
func invoice(r *gin.Engine, u repository.InvoiceRepo) {
	ih := handler.NewInvoiceHandler(u, lg)
	route := r.Group("/api/invoice")
	{
		route.GET("", ih.List())
	}
}
