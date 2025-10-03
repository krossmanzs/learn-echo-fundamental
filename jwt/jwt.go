package jwt

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Credential struct {
	Username string `json:"username"`
}

// Create a secret key
var jwtSecret = []byte("super-secret-key")

// generate a token (login endpoint)
func login(c echo.Context) error {
	// validate username/password
	username := "krossmanzs"

	// the username requested must be the same
	// so lets validate it
	cred := new(Credential)

	if err := c.Bind(cred); err != nil {
		return err
	} else if cred.Username != username {
		return c.String(http.StatusUnauthorized, "Username salah!")
	}


	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // expires in 1 hour
	})

	// Sign it
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

/* 
	a handler function for restricted route
	By default: "user" is where echo-jwt stashes the validated token.
	If you change it in your code without updating the middleware, you'll just get nil and crash.
	If you want a different key, configure it with ContextKey.
*/
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+username+"!")
}

func RunJwt() {
	e := echo.New()

	// public route
	e.POST("/login", login)

	// Restricted group
	r := e.Group("/restricted")

	// adding jwt middleware to the grouped routes
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtSecret,
	}))

	r.GET("", restricted) // now protected
	e.Logger.Fatal(e.Start(":8080"))
}