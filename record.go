package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Book struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Pic    string `db:"pic"`
}

func Pullbook(c echo.Context) error {
	var book Book
	title := c.Param("title")
	sess.Select("*").From(tablename).Where("title = ?", title).Load(&book)

	return c.JSON(http.StatusOK, book)
}

func Pullbooks(c echo.Context) error {
	var books []Book
	sess.Select("*").From(tablename).Load(&books)

	return c.JSON(http.StatusOK, books)
}

func Updatebook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": book.Id, "title": book.Title, "author": book.Author, "pic": book.Pic}
	sess.Update(tablename).SetMap(attrsMap).Where("id = ?", book.Id).Exec()

	return c.NoContent(http.StatusOK)
}

func Postbook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	sess.InsertInto(tablename).Columns("title", "author", "pic").Values(book.Title, book.Author, book.Pic).Exec()

	return c.NoContent(http.StatusOK)
}

func Deletebook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sess.DeleteFrom(tablename).Where("id = ?", id).Exec()

	return c.NoContent(http.StatusNoContent)
}
