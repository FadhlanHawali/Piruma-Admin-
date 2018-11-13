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


	if err:= idb.DB.Raw("SELECT * FROM orders WHERE(timestamp_start BETWEEN ? AND ?) OR (timestamp_end BETWEEN ? AND ?) OR ((timestamp_start <= ? AND timestamp_start <= ?) AND (timestamp_end >= ? AND timestamp_end >= ?))",timeStamp.TimestampStart,timeStamp.TimestampEnd,timeStamp.TimestampStart,timeStamp.TimestampEnd,timeStamp.TimestampStart,timeStamp.TimestampEnd,timeStamp.TimestampStart,timeStamp.TimestampEnd).Find(&order).Error;err!=nil{

	}else {

		result = gin.H{
			"status":"failed",
			"reason":"Jadwal Tabrakan",
			"jam":time.Now().Unix(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	//if(timeStamp.TimestampStart <=timestamp || timeStamp.TimestampEnd <= timestamp){
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"status":"failed",
	//		"reason":"Time has passed
	//	})
	//	return
	//}

	order.StatusSurat = statusSurat.StatusSurat
	order.StatusPeminjaman = statusSurat.StatusPeminjaman

	order.TimestampStart = timeStamp.TimestampStart
	order.TimestampEnd = timeStamp.TimestampEnd

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

func (idb *InDB) PublicSearch (c *gin.Context){
	var(
		search model.SearchRuangan
		//timestamp model.TimeStamp
		//order model.Orders
		arr [] model.Hasil
		result gin.H
	)

	if err:= c.Bind(&search);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:= idb.DB.Raw("SELECT COUNT(id_departemen),id_departemen FROM ruangans WHERE kapasitas >= ? GROUP BY id_departemen ",search.Kapasitas).Find(&arr).Error;err != nil{
		result = gin.H{
			"result":arr,
			"count":len(arr),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}else {

		result = gin.H{
			"result": arr,
			"count":  len(arr),
		}
		c.JSON(http.StatusOK, result)
		return
	}

}

func (idb *InDB) PublicListRoom (c *gin.Context){
	var(
		ruangan [] model.Ruangan
		result gin.H
	)

	idDepartemen := c.Param("idDepartemen")
	kapasitas := c.Query("kapasitas")

	if err := idb.DB.Raw("select * from ruangans where id_departemen = ? AND kapasitas >= ?",idDepartemen,kapasitas).Find(&ruangan).Error;
		err != nil{
		result = gin.H{
			"result":"Ruangan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		result = gin.H{
			"result":ruangan,
		}
		c.JSON(http.StatusOK,result)
	}
}

func (idb *InDB) PublicDetailSchedule (c *gin.Context){

	var(
		jadwal [] model.Jadwal
		result gin.H
	)


	idRuangan := c.Param("idRuangan")
	timestamp_start := c.Query("start")
	timestamp_end := c.Query("end")

	if err:= idb.DB.Raw("select timestamp_start,timestamp_end,keterangan from orders where timestamp_start >= ? AND timestamp_end <= ? AND id_ruangan = ?",timestamp_start,timestamp_end,idRuangan).Find(&jadwal).Error;err!=nil{
		result = gin.H{
			"result":jadwal,
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		result = gin.H{
			"result":jadwal,
		}
		c.JSON(http.StatusOK,result)
		return
	}

}






