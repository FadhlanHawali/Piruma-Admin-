package config

import (
	"github.com/jinzhu/gorm"
	"PirumaAdmin/model"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "admin:BPGHQYSKEQJVQNFG@tcp(sl-aus-syd-1-portal.5.dblayer.com:20314)/fakultas")

	//db, err := gorm.Open("mysql", "root:pintar123@/fakultas?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}

	db.AutoMigrate(model.Departemen{},model.Ruangan{},model.Orders{})
	return db
}
