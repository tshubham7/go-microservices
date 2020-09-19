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

	conn, err := grpc.Dial("172.17.0.1:5000", grpc.WithInsecure()) // local
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cc := protos.NewInvoiceClient(conn)

	r := gin.Default()

	ur := repository.NewUserRepo(wdb)
	lr := repository.NewLogRepo(wdb)

	user(r, ur, lr, cc)
	servicelogs(r, lr)

	r.Run(":9001")
}

// user routes
func user(r *gin.Engine, u repository.UserRepo, lr repository.LogRepo, cc protos.InvoiceClient) {
	s := handler.NewUserHandler(u, lr, l, cc)
	route := r.Group("/api/user")
	{
		route.POST("", s.Create())
		route.GET("", s.List())
		route.DELETE("/:id", s.Delete())
		route.PATCH("/:id", s.Update())
	}
}

// servicelogs routes
func servicelogs(r *gin.Engine, lr repository.LogRepo) {
	lh := handler.NewLogHandler(lr, l)
	route := r.Group("/api/user/service-logs")
	{
		route.GET("", lh.List())
	}
}
