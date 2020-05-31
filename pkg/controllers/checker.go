package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go-find-me/pkg/models"
)

var NewChecker models.Checker

func CreateChecker(ctx *fasthttp.RequestCtx) {
	c := &models.Checker{}

	if err := json.Unmarshal(ctx.PostBody(), &c); err == nil {
		c.CreateChecker()

		json.NewEncoder(ctx).Encode(c)
	}
}

func GetCheckers(ctx *fasthttp.RequestCtx) {
	cs := models.GetAllCheckers()
	json.NewEncoder(ctx).Encode(cs)
}

func GetCheckerById(ctx *fasthttp.RequestCtx) {
	idString := fmt.Sprint(ctx.UserValue("id"))
	id := uuid.MustParse(idString)
	cs := models.GetCheckerById(id)
	json.NewEncoder(ctx).Encode(cs)
}

func DeleteChecker(ctx *fasthttp.RequestCtx) {
	idString := fmt.Sprint(ctx.UserValue("id"))
	id := uuid.MustParse(idString)
	cs := models.DeleteChecker(id)
	json.NewEncoder(ctx).Encode(cs)
}