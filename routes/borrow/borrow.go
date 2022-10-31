package borrow

import (
	"time"

	"github.com/Agilistikmal/perpustakaan-api/database"
	"github.com/gin-gonic/gin"
)

type Borrow struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	BookID        string    `json:"book_id"`
	Username      string    `json:"username"`
	BorrowDate    time.Time `json:"release_date"`
	MaxReturnDate time.Time `json:"max_return_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

var db = database.Connect()
var borrow Borrow
var borrows = []Borrow{}

func GetBorrowListByUsername(c *gin.Context) {
	var username = c.Param("username")

	db.Find(&borrows, "username = ?", username)

	if len(borrows) == 0 {
		c.JSON(200, gin.H{
			"status_code": 404,
			"error":       "Data tidak ditemukan.",
			"data":        nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        borrows,
	})
}

func AddBorrowList(c *gin.Context) {
	var body Borrow

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	db.Find(&borrows, "username = ? AND book_id = ?", body.Username, body.BookID)
	if len(borrows) > 0 {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       "Username tersebut sudah meminjam book_id tersebut sebelumnya.",
			"data":        nil,
		})
		return
	}

	db.Create(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menambahkan data peminjaman buku.",
	})
}

func UpdateBorrowList(c *gin.Context) {
	var username = c.Param("username")
	var book_id = c.Param("book_id")
	var body Borrow

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	if body.ID != borrow.ID {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       "Anda tidak dapat mengubah ID saat ini.",
			"data":        nil,
		})
		return
	}

	db.Where("username = ?", username).Where("book_id = ?", book_id).Updates(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil mengubah data peminjaman buku.",
	})
}

func DeleteBorrowList(c *gin.Context) {
	username := c.Param("username")
	book_id := c.Param("book_id")

	db.Where("username", username).Where("book_id", book_id).Delete(&Borrow{})
	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menghapus data.",
	})
}
