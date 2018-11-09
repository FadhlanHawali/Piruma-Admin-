package main

import (
	"PirumaAdmin/config"
	"PirumaAdmin/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main(){
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()
	router.POST("/api/signup", inDB.SignUp)
	router.POST("/api/login",inDB.Login)
	router.POST("/api/ruangan/add",inDB.AddRuangan)
	router.GET("/api/ruangan/:idRuangan/detail",inDB.DetailRuangan)
	router.GET("/api/ruangan/:idRuangan/delete",inDB.DeleteRuangan)
	router.GET("/api/ruang/list",inDB.ListRuangan)
	//router.POST("/cobo",inDB.AddOrder)
	router.POST("/api/addOrder",inDB.AddOrder)
	router.Run(":8080")
}
