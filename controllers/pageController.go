package controllers

import "github.com/gin-gonic/gin"

func ShowStartPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func ShowMapPage(c *gin.Context) {
	c.HTML(200, "map.html", nil)
}
