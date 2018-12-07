package main

import (
	"PirumaAdmin/config"
	"PirumaAdmin/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"

	"PirumaAdmin/middleware"
	"net/http"
	"github.com/gin-contrib/cors"
)

func main(){
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/api/signup", inDB.SignUp)
	router.POST("/api/login",inDB.Login)
	router.POST("/api/ruangan/add",middleware.Auth,inDB.AddRuangan)
	router.GET("/api/ruangan/:idRuangan/detail",inDB.DetailRuangan)
	router.GET("/api/ruangan/:idRuangan/delete",inDB.DeleteRuangan)
	router.PUT("/api/ruangan/:idRuangan/update",inDB.UpdateRuangan)
	router.GET("/api/ruang/list",inDB.ListRuangan)
	router.POST("/api/search",inDB.PublicSearch)
	router.GET("api/search/:idDepartemen",inDB.PublicListRoom)
	router.POST("/api/addOrder",middleware.Auth,inDB.AddOrder)
	router.POST("/api/public/addOrder",middleware.PublicAuth,inDB.PublicAddOrder)
	router.GET("/api/ruangan/:idRuangan/time",inDB.PublicDetailSchedule)
	router.GET("/api/jadwal/:idDepartemen/time",inDB.JadwalDepartemen)
	router.PUT("/api/order/:idPemesanan/accept",inDB.AcceptOrder)
	router.PUT("/api/order/:idPemesanan/decline",inDB.DeclineOrder)
	router.GET("/api/order/listorder",inDB.ListPublicOrder)
	router.PUT("/api/order/:idPemesanan/update",inDB.UpdateStatusSurat)
	router.GET("/api/order/check",middleware.PublicAuth,inDB.PublicCekStatusOrder)


	router.Run(":8080")
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers,Authorization,Content-Type")
	c.JSON(http.StatusOK, struct{}{})
}

