package controllers

import (
"encoding/json"
"fmt"
"github.com/google/uuid"
"github.com/valyala/fasthttp"
"go-find-me/pkg/models"
)

var NewDevice models.Device

func CreateDevice(ctx *fasthttp.RequestCtx) {
	c := &models.Device{}

	if err := json.Unmarshal(ctx.PostBody(), &c); err == nil {
		c.CreateDevice()

		json.NewEncoder(ctx).Encode(c)
	}
}

func GetDevices(ctx *fasthttp.RequestCtx) {
	cs := models.GetAllDevices()
	json.NewEncoder(ctx).Encode(cs)
}

func GetDeviceById(ctx *fasthttp.RequestCtx) {
	idString := fmt.Sprint(ctx.UserValue("id"))
	id := uuid.MustParse(idString)
	cs := models.GetDeviceById(id)
	json.NewEncoder(ctx).Encode(cs)
}

func DeleteDevice(ctx *fasthttp.RequestCtx) {
	idString := fmt.Sprint(ctx.UserValue("id"))
	id := uuid.MustParse(idString)
	cs := models.DeleteDevice(id)
	json.NewEncoder(ctx).Encode(cs)
}