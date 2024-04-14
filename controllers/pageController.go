package controllers

import "github.com/gin-gonic/gin"

func ShowStartPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
