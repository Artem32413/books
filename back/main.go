package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct{
	name string `name: "name" binding: "required"`
	email string `form: "email" binding: "required, email"`
	password string `form: "password" binding: "required, password"`
}

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

var books = []book{
	{ID: "1", Title: "Законы человеческой природы", Author: "Роберт Грин", Genre: "Психология"},
	{ID: "2", Title: "Старикам тут не место", Author: "Кормака Маккартхи", Genre: "Роман"},
	{ID: "3", Title: "Так говорил Заратустра", Author: "Фридриха Ницше", Genre: "Философия"},
}

var tmpl *template.Template

func postHtml(c *gin.Context){
	c.HTML(http.StatusOK ,"enterReg.html", nil)
}

func heandleForm(c *gin.Context){
	var data user 
	if err := c.ShouldBind(&data); err != nil{
		c.HTML(http.StatusBadRequest, "enterReg.html", gin.H{"error": err.Error()} )
	return}
	c.HTML(http.StatusOK ,"enterReg.html", gin.H{"name": data.name, "email": data.email, "password": data.password})
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, el := range books {
		if el.ID == id {
			c.IndentedJSON(http.StatusOK, el)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "books not foun"})
}

func deleteById(c *gin.Context) {
	id := c.Param(":id")
	for _, el := range books {
		if el.ID == id {
			c.IndentedJSON(http.StatusOK, el)
			continue
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "удаленныйое сообещние"})
}

func ginE() {
	router := gin.Default()
	router.LoadHTMLFiles(
		"./index.html",
		"./navigation/book/book.html",
		"./navigation/profile/profile.html",
		"./navigation/contacts/contacts.html",
		"./navigation/enterReg/enterReg.html",
	)
	// router.LoadHTMLGlob("./navigation/book/book.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil) // Отправка HTML на запрос
	})
	router.GET("/book/", func(c *gin.Context) {
		c.HTML(200, "book.html", nil) // Отправка HTML на запрос
	})
	router.GET("/profile/", func(c *gin.Context) {
		c.HTML(200, "profile.html", nil) // Отправка HTML на запрос
	})
	router.GET("/contacts/", func(c *gin.Context) {
		c.HTML(200, "contacts.html", nil) // Отправка HTML на запрос
	})
	router.GET("/enterReg/", func(c *gin.Context) {
		c.HTML(200, "enterReg.html", nil) // Отправка HTML на запрос
	})
	router.Static("/styles", "./styles")
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.GET("/books/:id/deleted", deleteById)
	router.POST("/books", postBook)
	router.Run("localhost:8080")
}

func main() {
	ginE()
}
