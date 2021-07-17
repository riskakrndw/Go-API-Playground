package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

func main() {
	// connection string user:password@tcp(hostname:port)/db
	db, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/alta")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Masukkan pilihan (lihat/tambah)")
	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "lihat":
		row, err := db.Query("SELECT id, title, author from books") // select (tidak mengubah isi tabel)
		if err != nil {
			panic(err.Error())
		}
		for row.Next() {
			var book Book
			err = row.Scan(&book.Id, &book.Title, &book.Author)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("Title:", book.Title)
			fmt.Println("Author:", book.Author)
		}
	case "tambah":
		newBook := Book{}
		fmt.Println("Title")
		fmt.Scanln(&newBook.Title)
		fmt.Println("Author")
		fmt.Scanln(&newBook.Author)
		_, err := db.Exec(fmt.Sprintf("INSERT INTO books(title, author) VALUES ('%s', '%s')", newBook.Title, newBook.Author))
		if err != nil {
			panic(err.Error())
		}
	}
}