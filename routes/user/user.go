package user

import (
	"time"

	"github.com/Agilistikmal/perpustakaan-api/database"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	FullName  string    `json:"release_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db = database.Connect()
var users = []User{}
var user User

func GetUsers(c *gin.Context) {
	db.Find(&users)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        users,
	})
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	db.First(&user, "username = ?", username)

	if user.Username == "" {
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
		"data":        users,
	})
}

func AddUser(c *gin.Context) {
	var body User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	db.Find(&users, "username = ?", body.Username)
	if len(users) > 0 {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       "Username tersebut sudah terdaftar sebelumnya.",
			"data":        nil,
		})
		return
	}

	db.Create(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menambahkan data user.",
	})
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var body User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"error":       err.Error(),
			"data":        nil,
		})
		return
	}

	if body.Username == "" {
		c.JSON(200, gin.H{
			"status_code": 404,
			"error":       "Data tidak ditemukan.",
			"data":        nil,
		})
		return
	}

	db.Where("username = ?", username).Updates(&body)

	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil mengubah data.",
	})
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")

	db.Where("username = ?", username).Delete(user)
	c.JSON(200, gin.H{
		"status_code": 200,
		"error":       nil,
		"data":        "Berhasil menghapus data.",
	})
}
