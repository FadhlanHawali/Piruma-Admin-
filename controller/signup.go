package controller

import (
	"github.com/gin-gonic/gin"
	"PirumaAdmin/model"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) SignUp(c *gin.Context)  {

	var(
		departemen model.Departemen
		signup model.SignUp
		result gin.H
	)

	if err:= c.Bind(&signup); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	if err:= idb.DB.Where(&model.Departemen{Username:signup.Username}).First(&departemen).Error;err != nil{
	}else {
		result = gin.H{
			"status":"failed",
			"reason":"Username does exist",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	departemen.IdDepartemen = "Dept"+"-"+ signup.IdDepartemen
	departemen.Departemen = signup.Departemen
	departemen.Username = signup.Username
	departemen.Email = signup.Email
	departemen.Fakultas = signup.Fakultas
	departemen.Kontak = signup.Kontak

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost);if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	departemen.Password = string(hashedPassword)

	idb.DB.Create(&departemen)
	result = gin.H{
		"status":"success",
	}
	c.JSON(http.StatusOK, result)

}