package routes

import (
"github.com/fasthttp/router"
"go-find-me/pkg/controllers"
)

var DeviceRoutes = func(r *router.Router) {
	r.GET("/devices", controllers.GetDevices)
	r.GET("/devices/{id}", controllers.GetDeviceById)
	r.POST("/devices", controllers.CreateDevice)
	r.DELETE("/devices/{id}", controllers.DeleteDevice)
}
