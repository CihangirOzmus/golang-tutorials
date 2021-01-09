package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	a "main/animal"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func getHome(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Hello, Go Echo!</strong>")
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			return c.JSON(http.StatusOK, u)
		}
	}

	return c.JSON(http.StatusNotFound, "User does not exist!")
}

func saveUser(c echo.Context) (err error) {
	u := new(User)
	uId, _ := uuid.NewRandom()
	u.ID = uId.String()
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	users = append(users, User{ID: u.ID, Name: u.Name, Email: u.Email})
	return c.JSON(http.StatusCreated, u)
}

func updateUser(c echo.Context) (err error) {
	id := c.Param("id")
	updateUserRequest := echo.Map{}

	if err := c.Bind(&updateUserRequest); err != nil {
		return err
	}

	for index, u := range users {
		if u.ID == id {
			users[index].Name = fmt.Sprint(updateUserRequest["name"])
			users[index].Email = fmt.Sprint(updateUserRequest["email"])
			return c.JSON(http.StatusOK, "User is updated.")
		}
	}

	return echo.NewHTTPError(http.StatusBadRequest, "User could not be updated!")
}

func deleteUser(c echo.Context) (err error) {
	id := c.Param("id")

	for index, u := range users {
		if u.ID == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, "User is deleted.")
		}
	}

	return echo.NewHTTPError(http.StatusBadRequest, "User does not found!")
}

func main() {
	users = append(users, User{ID: "1", Name: "Cihangir", Email: "cozmus@gmail.com"})
	users = append(users, User{ID: "2", Name: "Dummy", Email: "dummy@gmail.com"})

	e := echo.New()
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", getHome)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.POST("/cats", a.AddCat)
	e.POST("/dogs", a.AddDog)
	e.POST("/hamsters", a.AddHamster)
	e.Logger.Fatal(e.Start(":8080"))
}
