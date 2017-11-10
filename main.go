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
	e.GET("/book/api/fetch/:title", PullBookInfo)
	e.GET("/book/api/fetch/all", PullBooksInfo)
	e.PUT("/book/api/update", UpdateBookInfo)
	e.POST("/book/api/post", PostBookInfo)
	e.DELETE("/book/api/delete/:id", DeleteBookInfo)

	e.Logger.Fatal(e.Start(":9090"))
}
