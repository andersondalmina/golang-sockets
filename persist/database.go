package persist

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func CreateDatabase() {
	db = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		fmt.Println("Unable to connecto to database")
		os.Exit(1)
	}
}

func GetBookByTitle(title string) (Book, error) {
	var book Book
	err := db.Model(&book).
		Where("titulo LIKE ?", title).
		Relation("Editions").
		Limit(1).
		Select()

	return book, err
}

func GetBookByID(id int64) (Book, error) {
	var book Book
	err := db.Model(&book).
		Where("codigo = ?", id).
		Limit(1).
		Select()

	return book, err
}

func CreateBook(title string, author string, edition int64, year int64) (Book, error) {
	book := Book{
		ID:      12345,
		Title:   title,
		Author:  author,
		Edition: edition,
		Year:    year,
	}

	_, err := db.Model(&book).Insert()
	if err != nil {
		panic(err)
	}

	return book, err
}

func UpdateBook(id int64, title string, author string, edition int64, year int64) (Book, error) {
	book, err := GetBookByID(id)

	book.Title = title
	book.Author = author
	book.Edition = edition
	book.Year = year

	_, err = db.Model(&book).Where("codigo = ?", book.ID).Update()
	if err != nil {
		panic(err)
	}

	return book, err
}

func SearchBookByTitle(term string) []Book {
	var books []Book
	err := db.Model(&books).
		Where("titulo LIKE ?", "%"+term+"%").
		Order("titulo ASC").
		Select()

	if err != nil {
		panic(err)
	}

	return books
}

func SearchBooksByEdition(edition string) []Book {
	var books []Book
	err := db.Model(&books).
		Where("edicao = ?", edition).
		Order("titulo ASC").
		Select()

	if err != nil {
		panic(err)
	}

	return books
}

func SearchBookByAuthor(author string) []Book {
	var books []Book
	err := db.Model(&books).
		Where("autor LIKE ?", "%"+author+"%").
		Order("titulo ASC").
		Select()

	if err != nil {
		panic(err)
	}

	return books
}

func SearchBookByYear(year string) []Book {
	var books []Book
	err := db.Model(&books).
		Where("ano = ?", year).
		Order("titulo ASC").
		Select()

	if err != nil {
		panic(err)
	}

	return books
}

func DeleteBooksByTitle(title string) error {
	var book Book
	_, err := db.Model(&book).
		Where("TRIM(titulo) = ?", title).
		Delete()

	return err
}
