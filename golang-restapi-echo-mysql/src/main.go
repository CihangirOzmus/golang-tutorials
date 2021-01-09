package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Username  string    `json:"nickname"`
	Email     string    `json:"email"`
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get All Users")
}

func getUserById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, fmt.Sprintf("Get User by id %s", id))
}

func postUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Post User")
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, fmt.Sprintf("Update User %s", id))
}

func deleteUserById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, fmt.Sprintf("Delete User by id %s", id))
}

func main()  {
	fmt.Println("Server is started ", time.Now())
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/", mainHandler)

	users := e.Group("/users")
	users.GET("/", getUsers)
	users.GET("/:id", getUserById)
	users.POST("/", postUser)
	users.PUT("/:id", updateUser)
	users.DELETE("/:id", deleteUserById)


	dsn := "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("No database connection!")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Username: "cozmus", Email: "cozmus@gmail.com"})

	// Read
	var user User
	db.First(&user, 1)
	db.First(&user, "username = ?", "cozmus")

	// Update - update product's price to 200
	db.Model(&user).Update("Email", "updated@gmail.com")
	// Update - update multiple fields
	db.Model(&user).Updates(User{Email: "updated2@gmail.com", Username: "updated2"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"Email": "updated3@gmail.com", "Username": "updated3"})

	// Delete - delete product
	db.Delete(&user, 1)

	e.Logger.Fatal(e.Start(":8080"))
}
