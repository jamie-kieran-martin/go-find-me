package controllers

import (
	"encoding/json"
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