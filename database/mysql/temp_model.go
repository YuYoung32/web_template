package mysql

import (
	"gorm.io/gorm"
	"web_template/log"
)

type TempModel struct {
	Name   string `json:"name" gorm:"type:varchar(255) COLLATE utf8_bin"`
	IP     string `json:"ip"`
	Region string `json:"region"`
	Time   int64  `json:"time" gorm:"autoCreateTime"`
}

func autoMigrateTempModel(db *gorm.DB) {
	err := db.AutoMigrate(&TempModel{})
	if err != nil {
		log.GetLogger().Error("auto migrate failed: " + err.Error())
	}
}
