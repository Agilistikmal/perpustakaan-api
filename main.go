package main

import (
	"github.com/Agilistikmal/perpustakaan-api/routes/books"
	"github.com/Agilistikmal/perpustakaan-api/routes/borrow"
	"github.com/Agilistikmal/perpustakaan-api/routes/user"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	/*
	* 	Book Router
	 */
	app.GET("/book", books.GetBooks)
	app.GET("/book/:book_id", books.GetBookById)
	app.POST("/book", books.AddBook)
	app.PUT("/book/:book_id", books.UpdateBook)
	app.DELETE("/book/:book_id", books.DeleteBook)

	/*
	*		Borrow Router
	 */
	app.GET("/borrow/:username", borrow.GetBorrowListByUsername)
	app.POST("/borrow", borrow.AddBorrowList)
	app.PUT("/borrow/:username/:book_id", borrow.UpdateBorrowList)
	app.DELETE("/borrow/:username/:book_id", borrow.DeleteBorrowList)

	/*
	*		User Router
	 */
	app.GET("/user", user.GetUsers)
	app.GET("/user/:username", user.GetUserByUsername)
	app.POST("/user", user.AddUser)
	app.PUT("/user/:username", user.UpdateUser)
	app.DELETE("/user/:username", user.DeleteUser)

	/*
	*		Invalid Routes
	 */
	app.LoadHTMLFiles("./templates/404.html")
	app.NoRoute(func(c *gin.Context) {
		c.HTML(200, "404.html", gin.H{"ip": c.ClientIP()})
	})

	app.Run("127.0.0.1:8080")
}
