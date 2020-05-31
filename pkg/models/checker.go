package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go-find-me/pkg/config"
	"time"
)

var db *gorm.DB

type Checker struct {
	Id uuid.UUID `gorm:"primary_key"json:"id"`
	DeviceId uuid.UUID `gorm:"TYPE:blob REFERENCES devices"`
	Time time.Time `json:"time"`
	Result bool `json:"result"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Checker{})
	db.Model(&Checker{})
}

func (c *Checker) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.New())
	return
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

func GetCheckerById(Id uuid.UUID) *Checker {
	var c Checker
	db.Where("ID = ?", Id).Find(&c)
	return &c
}

func DeleteChecker(Id uuid.UUID) Checker {
	var c Checker
	db.Where("ID = ?", Id).Delete(c)
	return c
}
