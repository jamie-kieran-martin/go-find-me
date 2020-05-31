package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go-find-me/pkg/config"
)

type Device struct {
	Id uuid.UUID `gorm:"primary_key"json:"id"`
	Name string
	Address string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Device{})
}

func (c *Device) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func (c *Device) CreateDevice() *Device {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllDevices() []Device {
	var cs []Device
	db.Find(&cs)
	return cs
}

func GetDeviceById(Id uuid.UUID) *Device {
	var c Device
	db.Where("ID = ?", Id).Find(&c)
	return &c
}

func DeleteDevice(Id uuid.UUID) Device {
	var c Device
	db.Where("ID = ?", Id).Delete(&c)
	return c
}
