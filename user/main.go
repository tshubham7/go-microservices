package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/user/db"
	"github.com/tshubham7/go-microservices/user/handler"
	"github.com/tshubham7/go-microservices/user/repository"
	"google.golang.org/grpc"
)

var l = log.New(os.Stdout, "user-service ", log.LstdFlags)

func main() {
	wdb, err := db.GetDatabase()
	if err != nil {
		l.Fatal(err)
	}

	db.RunMigrations(wdb)
	l.Println("successfully migrated models")

	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure()) // local
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cc := protos.NewInvoiceClient(conn)

	r := gin.Default()

	ur := repository.NewUserRepo(wdb)
	user(r, ur, cc)

	r.Run(":9001")
}

func user(r *gin.Engine, u repository.UserRepo, cc protos.InvoiceClient) {
	s := handler.NewUserHandler(u, l, cc)
	route := r.Group("/api/user")
	{
		route.POST("/", s.Create())
		route.GET("/", s.List())
		route.DELETE("/:id", s.Delete())
		route.PATCH("/:id", s.Update())
	}

}
