package controller

import (
	"github.com/gin-gonic/gin"
	"PirumaAdmin/model"
	"net/http"
	"strconv"
	"time"
)

//func (idb *InDB) AddOrder (c *gin.Context){
//	var(
//		cobo model.Cobo
//		objek model.ObjectBanyak
//	)
//
//	if err:= c.Bind(&cobo);err != nil{
//		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
//		return
//	}
//
//	objek = cobo.Objek
//	c.JSON(http.StatusOK,gin.H{
//		"cobo1" : cobo.Nama,
//		"result": objek.Objek1,
//	})
//	return
//
//}

func (idb *InDB) AddOrder (c *gin.Context){
	var(
		addOrder model.AddOrder
		timeStamp model.TimeStamp
		statusSurat model.StatusSurat
		order model.Orders
		ruangan model.Ruangan
		departemen model.Departemen

		result gin.H
	)

	if err:= c.Bind(&addOrder);err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	statusSurat = addOrder.StatusSurat
	timeStamp = addOrder.TimeStamp

	if err:= idb.DB.Where(&model.Ruangan{IdRuangan:addOrder.IdRuangan,IdDepartemen:addOrder.IdDepartemen}).First(&ruangan).Error;err != nil{
		result = gin.H{
			"status":"failed",
			"reason":"Ruangan doesn't exist",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	if err:= idb.DB.Where(&model.Departemen{IdDepartemen:addOrder.IdDepartemen}).First(&departemen).Error;err != nil{
		result = gin.H{
			"status":"failed",
			"reason":"Departemen doesn't exist",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	timestamp := strconv.FormatInt(time.Now().Unix(),10)

	//if(timeStamp.TimestampStart <=timestamp || timeStamp.TimestampEnd <= timestamp){
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"status":"failed",
	//		"reason":"Time has passed",
	//	})
	//	return
	//}

	order.StatusSurat = statusSurat.StatusSurat
	order.StatusPeminjaman = statusSurat.StatusPeminjaman

	order.TimestampStart = timeStamp.TimestampStart
	order.TimestapEnd = timeStamp.TimestampEnd

	order.Ruangan = addOrder.Ruangan
	order.IdDepartemen = addOrder.IdDepartemen
	order.IdRuangan = addOrder.IdRuangan
	order.Departemen = addOrder.Departemen
	order.Email = addOrder.Email

	order.IdPemesanan = "rent"+"-"+string(timestamp)
	order.PenanggungJawab = addOrder.PenanggungJawab
	order.Keterangan = addOrder.Keterangan
	order.Telepon = addOrder.Telepon

	idb.DB.Create(&order)

	result = gin.H{
		"status":"success",
	}
	c.JSON(http.StatusOK, result)

	return

}

func (idb *InDB) PublicOrder (c *gin.Context){

}



