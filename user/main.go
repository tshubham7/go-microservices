package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/go-microservices/user/db"
	"github.com/tshubham7/go-microservices/user/handler"
	"github.com/tshubham7/go-microservices/user/repository"
)

var l = log.New(os.Stdout, "user-service ", log.LstdFlags)

func main() {
	wdb, err := db.GetDatabase()
	if err != nil {
		l.Fatal(err)
	}

	db.RunMigrations(wdb)
	l.Println("successfully migrated models")

	r := gin.Default()

	ur := repository.NewUserRepo(wdb)
	user(r, ur)

	r.Run(":9001")
}

func user(r *gin.Engine, u repository.UserRepo) {
	s := handler.NewUserHandler(u, l)
	route := r.Group("/api/user")
	{
		route.POST("/", s.Create())
		route.GET("/", s.List())
		route.DELETE("/:id", s.Delete())
		route.PATCH("/:id", s.Update())
	}

}
