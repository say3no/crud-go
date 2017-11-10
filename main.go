package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

var (
	tablename = "test"
	conn, _   = dbr.Open("mysql", "username:api@/book_db", nil)
	sess      = conn.NewSession(nil)
)

func main() {
	e := echo.New()

	// Routing
	e.GET("/book/api/fetch/:title", Pullbook)
	e.GET("/book/api/fetch/all", Pullbooks)
	e.PUT("/book/api/update", Updatebook)
	e.POST("/book/api/post", Postbook)
	e.DELETE("/book/api/delete/:id", Deletebook)

	e.Logger.Fatal(e.Start(":9090"))
}
