// Package status is here just to say that the API is up and that it can
// access the CouchDB databases, for debugging and monitoring purposes.
package status

import (
	"github.com/gin-gonic/gin"
)

// Status responds OK if the service is running
func Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// Routes sets the routing for the status service
func Routes(router *gin.RouterGroup) {
	router.GET("/", Status)
}
