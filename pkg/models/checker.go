package models

import (
	"go-find-me/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

type Checker struct {
	gorm.Model
	// Id int64 `gorm:""json:"id"`
	Time time.Time `json:"time"`
	Result bool `json:"result"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Checker{})
}

func (c *Checker) CreateChecker() *Checker {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCheckers() []Checker {
	var cs []Checker
	db.Find(&cs)
	return cs
}

func GetCheckerById(Id int64) (*Checker, *gorm.DB) {
	var c Checker
	db := db.Where("ID = ?", Id).Find(&c)
	return &c, db
}

func DeleteChecker(Id int64) Checker {
	var c Checker
	db.Where("ID = ?", Id).Delete(c)
	return c
}
