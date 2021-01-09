package animal

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to read request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed to unmarshalling in addCat %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is new cat %#v", cat)
	return c.String(http.StatusOK, "We got your cat.")
}

func AddDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)

	if err != nil {
		log.Printf("Failed to process addDog %s", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	log.Printf("This is your dog %#v", dog)
	return c.String(http.StatusOK, "we got your dog.")
}

func AddHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed to process addHamster %s", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	log.Printf("This is your hamster %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster.")
}
