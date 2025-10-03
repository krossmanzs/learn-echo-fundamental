package basic

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	// Name string is Go field of type string
	// json:"name" is a struct tag. It tells Go's JSON encoder/decoder
	Name string `json:"name"`
	Email string `json:"email"`
}

var users []User

func RunBasic() {
	e := echo.New()

	// func(c echo.Context) error... is a handler (anonymous function) 
	// it returns error type, coz golang hate exceptions.
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hellom Echo!")
	})

	e.GET("/users", func(c echo.Context) error {
		if len(users) <= 0 {
			return c.String(http.StatusOK, "List user kosong!")
		}

		// convert collection into array of byte
		b, _ := json.Marshal(users)

		// return converted bytes into string
		return c.String(http.StatusOK, string(b))
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "User ID: " + id)
	})

	e.POST("/users", func (c echo.Context) error {
		u := new(User)

		/* 
			this is like a validation, do the body structure
			is correct like User struct or not
		*/
		if err := c.Bind(u); err != nil {
			return err
		}

		users = append(users, User{u.Name, u.Email})

		return c.JSON(http.StatusCreated, "Success")
	})

	e.Start(":8080")
}