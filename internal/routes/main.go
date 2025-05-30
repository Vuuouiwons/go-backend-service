package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run the server
func Run() {
	getRoutes()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addAblumsRoute(v1)

}
