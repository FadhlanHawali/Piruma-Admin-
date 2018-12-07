package controller

import (
	"github.com/gin-gonic/gin"
	"PirumaAdmin/model"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"fmt"
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

	ruangan.IdDepartemen = c.MustGet("id").(string)
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
	idb.DB.LogMode(true)

	idRuangan := c.Param("idRuangan")
	if err := idb.DB.Where("id_ruangan = ?",idRuangan).First(&ruangan).Error;
		err != nil {
		result = gin.H{
			"result": "Ruangan tidak ada",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	if err := idb.DB.Where("id_ruangan = ?",idRuangan).Delete(&ruangan).Error; err != nil{
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
		return
	}
}

func (idb *InDB) UpdateRuangan (c *gin.Context) {

	type updateRuangan struct {
		Kapasitas string `json:"kapasitas"`
		NamaRuangan string `json:"nama_ruangan"`
		Fasilitas string `json:"fasilitas"`
	}
	var (
		ruangan model.Ruangan
		update updateRuangan
		//result  gin.H
	)

	idRuangan := c.Param("idRuangan")
	if err:= c.Bind(&update);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	idb.DB.LogMode(true)

	//if err:=idb.DB.Raw("select * from ruangans where id_ruangan = ?",idRuangan).Find(&ruangan).Error;err!=nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"status":"failed",
	//		"reason":"Ruangan Tidak Ada",
	//	})
	//	return
	//}



	idb.DB.LogMode(true)
	if err:= idb.DB.Raw("select * from ruangans where id_ruangan = ?",idRuangan).Find(&ruangan).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"result":"Ruangan tidak ada",
		})
		return
	}

	if err:=idb.DB.Table("ruangans").Where("id_ruangan = ?",idRuangan).UpdateColumn("kapasitas", update.Kapasitas).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:=idb.DB.Table("ruangans").Where("id_ruangan = ?",idRuangan).UpdateColumn("nama_ruangan", update.NamaRuangan).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if err:=idb.DB.Table("ruangans").Where("id_ruangan = ?",idRuangan).UpdateColumn("fasilitas", update.Fasilitas).Error;err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}else {
		c.JSON(http.StatusOK,gin.H{
			"result":"success",
		})

		return
	}


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

func (idb *InDB) JadwalDepartemen (c *gin.Context){

	var (
		orders []model.Orders
		//result gin.H
	)
	idDepartemen := c.Param("idDepartemen")
	timestamp_start := c.Query("start")
	timestamp_end := c.Query("end")

	idb.DB.LogMode(true)
	if err:=idb.DB.Raw("select * from orders where status_proses = ? AND id_departemen = ? AND timestamp_start >= ? AND timestamp_end <= ?",true,idDepartemen,timestamp_start,timestamp_end).Find(&orders);err!=nil{

		var inInterface []map[string]interface{}
		inrec, _ := json.Marshal(orders)
		json.Unmarshal(inrec, &inInterface)

		gropped:= groupBy(inInterface,"ruangan")

		//jsonString, err := json.Marshal(gropped)
		//fmt.Println(err)

		c.JSON(http.StatusOK,gin.H{
			"result":gropped,
		})
		return
	}
	//
	//message := []map[string]interface{}{
	//	{"sda":"asd"},
	//}
}

type Hasil struct {
	Ruangan string
	Jadwal []map[string]interface{}
}

func groupBy(maps []map[string]interface{}, key string) []Hasil {
	//result := []
	groups := make(map[string][]map[string]interface{})
	for _,m := range maps {
		k := m[key].(string) // XXX: will panic if m[key] is not a string.
		groups[k] = append(groups[k], m)
		//fmt.Println(m[key].(string))
		//fmt.Println(m)
	}
	fmt.Println(len(groups))
	res := make([]Hasil,0)
	for i, _ := range groups {
		res = append(res, Hasil{
			Ruangan: i,
			Jadwal: groups[i],
		})
		fmt.Println(i)
	}

	return res
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

