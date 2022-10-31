package books

import (
	"time"

	"github.com/Agilistikmal/perpustakaan-api/database"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	BookID      string    `json:"book_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Writer      string    `json:"writer"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var db = database.Connect()
var books = []Book{}
var book Book

func GetBooks(c *gin.Context) {
	db.Find(&books)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        books,
	})
}

func GetBookById(c *gin.Context) {
	book_id := c.Param("book_id")
	db.First(&book, "book_id = ?", book_id)

	if book.BookID == "" {
		c.JSON(404, gin.H{
			"status_code": 404,
			"error":       "Buku dengan id tersebut tidak ditemukan",
			"data":        nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        book,
	})
}

func AddBook(c *gin.Context) {
	var body Book

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	db.Create(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menambahkan daftar buku.",
	})
}

func UpdateBook(c *gin.Context) {
	book_id := c.Param("book_id")
	var body Book

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	if body.BookID != book.BookID {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       "Anda tidak dapat mengubah ID saat ini.",
			"data":        nil,
		})
		return
	}

	db.Where("book_id = ?", book_id).Updates(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil memperbarui data.",
	})
}

func DeleteBook(c *gin.Context) {
	book_id := c.Param("book_id")

	db.Delete(Book{BookID: book_id})
	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menghapus data.",
	})
}
