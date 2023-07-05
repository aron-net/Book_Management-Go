package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/aron-helu/Book_Mangement-Go/pkg/controllers"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *httprouter.Router) {
	router.POST("/book", controllers.CreatBook)
	router.GET("/book", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
}
