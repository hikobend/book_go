package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Book_JSON struct { // JSON
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Page        int    `json:"page"`
	Description string `json:"description"`
}

type Book struct { // DB
	Id          int
	Title       string
	SubTitle    string
	Author      string
	Publisher   string
	Page        int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	r := gin.Default()
	r.POST("/create", insert)

	r.Run()
}

func insert(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var book Book_JSON
	c.ShouldBindJSON(&book)

	insert, err := db.Prepare("INSERT INTO book_go(title, sub_title, author, publisher, page, description) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	insert.Exec(book.Title, book.SubTitle, book.Author, book.Publisher, book.Page, book.Description)
}
