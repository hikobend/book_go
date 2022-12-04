package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
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
	r.POST("/create", Insert)
	r.GET("/books", Gets)
	r.GET("/book/:id", GetById)
	r.DELETE("/book/:id", Delete)

	r.Run()
}

func Insert(c *gin.Context) {
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

func Gets(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, title, sub_title, author, publisher, page, description, created_at, updated_at from book_go")
	if err != nil {
		log.Fatal(err)
	}
	var resultBook []Book
	for rows.Next() {
		book := Book{}
		if err := rows.Scan(&book.Id, &book.Title, &book.SubTitle, &book.Author, &book.Publisher, &book.Page, &book.Description, &book.CreatedAt, &book.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		resultBook = append(resultBook, book)
	}

	c.JSON(http.StatusOK, resultBook)
}

func GetById(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	var getbook Book
	if err = db.QueryRow("SELECT * FROM book_go where id = ?", id).Scan(&getbook.Id, &getbook.Title, &getbook.SubTitle, &getbook.Author, &getbook.Publisher, &getbook.Page, &getbook.Description, &getbook.CreatedAt, &getbook.UpdatedAt); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, getbook)
}

func Delete(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	delete, err := db.Prepare("DELETE FROM book_go WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	delete.Exec(id)
}
