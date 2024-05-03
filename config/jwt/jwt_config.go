package jwt

import (
	"os"

	"github.com/Clareand/customer_oder_synapsis/config/postgresql"
	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomer struct {
	Customer        Customer `json:"customer"`
	RefreshToken    string   `json:"refresh_token"`
	BussinessEntity string   `json:"id_business_entity"`
	jwt.StandardClaims
}

type Customer struct {
	CustomerID string `json:"customerID"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	CreatedAt  string `json:"createdAt"`
}

type GetCustomer struct {
	CustomerID string `json:"customerID"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	CreatedAt  string `json:"createdAt"`
}

func JWTConfig() middleware.JWTConfig {
	pubKey := []byte(os.Getenv("SECRET_KEY"))

	config := middleware.JWTConfig{
		SigningKey:    pubKey,
		SigningMethod: jwt.SigningMethodHS256.Name,
		Claims:        &JwtCustomer{},
	}
	return config
}

func Authz(apiName string, dbConn *postgresql.DbConnection) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}

func GetDataCustomer(c echo.Context) *GetCustomer {
	jwtUser := c.Get("user").(*jwt.Token)
	jwtClaims := jwtUser.Claims.(*JwtCustomer)
	jwtGetUsers := jwtClaims.Customer

	data := GetCustomer{
		CustomerID: jwtGetUsers.CustomerID,
		Username:   jwtGetUsers.Username,
		Email:      jwtGetUsers.Email,
		CreatedAt:  jwtGetUsers.CreatedAt,
	}

	return &data
}
