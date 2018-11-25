package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func PublicAuth(c *gin.Context) {

	var idMahasiswa string

	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		mapstructure.Decode(claims["id_mahasiswa"], &idMahasiswa)
		c.Set("id", idMahasiswa)
	} else {
		result := gin.H{
			"token":tokenString,
			"message": "token tidak valid",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}