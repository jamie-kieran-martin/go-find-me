package routes


import (
	"github.com/fasthttp/router"
	"go-find-me/pkg/controllers"
)

var CheckerRoutes = func(r *router.Router) {
	r.GET("/checkers", controllers.GetCheckers)
	r.POST("/checkers", controllers.CreateChecker)
}