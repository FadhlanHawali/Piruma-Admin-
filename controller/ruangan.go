package controller

import (
	"github.com/gin-gonic/gin"
	"PirumaAdmin/model"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDB) AddRuangan (c *gin.Context){

	var (
		ruangan    model.Ruangan
		addRuangan model.AddRuangan
		result     gin.H
	)

	if err:= c.Bind(&addRuangan); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	ruangan.IdDepartemen = addRuangan.IdDepartemen
	ruangan.Kapasitas = addRuangan.Kapasitas
	ruangan.NamaRuangan = addRuangan.NamaRuangan
	ruangan.Fasilitas = addRuangan.Fasilitas

	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	ruangan.IdRuangan = "room"+"-"+string(timestamp)

	idb.DB.Create(&ruangan)
	result = gin.H{
		"status":"success",
	}

	c.JSON(http.StatusOK,result)

	return
}

func (idb *InDB) DetailRuangan (c *gin.Context) {
	var(
		ruangan model.Ruangan
		result gin.H
	)

	idRuangan := c.Param("idRuangan")

	if err := idb.DB.Where("id_ruangan = ?",idRuangan).First(&ruangan).Error;
		err != nil{
		result = gin.H{
			"result":"Ruangan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		c.JSON(http.StatusOK,ruangan)
	}

}
func (idb *InDB) DeleteRuangan (c *gin.Context){
	var(
		ruangan model.Ruangan
		result gin.H
	)

	idRuangan := c.Param("idRuangan")

	if err := idb.DB.Where("id_ruangan = ?",idRuangan).Delete(&ruangan).Error;
		err != nil{
		result = gin.H{
			"result":"Ruangan tidak ada",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		result = gin.H{
			"result":"success",
		}
		c.JSON(http.StatusOK,result)
	}
}

func (idb *InDB) UpdateRuangan (c *gin.Context){

}

func (idb *InDB) ListRuangan (c *gin.Context){
	var(
		ruangan [] model.Ruangan
		result gin.H
	)

	if err := idb.DB.Where("id_departemen =?","Dept-1").Find(&ruangan).Error;
		err != nil{
		result = gin.H{
			"result":"Belum punya ruangan",
		}
		c.JSON(http.StatusBadRequest,result)
		return
	}else {
		result = gin.H{
			"result":ruangan ,
			"count":  len(ruangan),
		}
		c.JSON(http.StatusOK,result)
		return
	}
}




//func (idb *InDB) SearchRuangan (c *gin.Context){
//	var(yy
//		search model.SearchRuangan
//		ruangan model.Ruangan
//	)
//
//	if err:= c.Bind(&search); err!= nil{
//		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
//		return
//	}
//
//	time = search.TimeStamp
//
//	if err:= idb.DB.Where(&model.Orders{})
//}

