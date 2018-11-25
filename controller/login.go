package controller

import (
	"github.com/gin-gonic/gin"
	"PirumaAdmin/model"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)
func (idb *InDB) Login(c *gin.Context){
	var (
		login model.Login
		departemen model.Departemen
		result gin.H
	)

	if err:= c.Bind(&login); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:= idb.DB.Where(&model.Departemen{Username:login.Username}).First(&departemen).Error;err!=nil{
		result = gin.H{
			"result":"Username doesn't exist",
			"email":err,
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		var password_tes = bcrypt.CompareHashAndPassword([]byte(departemen.Password), []byte(login.Password))
		if password_tes == nil {
			sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username":departemen.Username,
				"id_departemen":departemen.IdDepartemen,
			})
			token, err := sign.SignedString([]byte("secret"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			//login success
			result = gin.H{
				"result":"Success",
				"token":token,
			}
			c.JSON(http.StatusOK,result)
			return
		} else {
			//login failed
			result = gin.H{
				"result":"Salah bos q",
			}
			c.JSON(http.StatusBadRequest,result)
			return
		}
	}
}

