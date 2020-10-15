package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", login)
	router.POST("/report", report)
	router.POST("/headset", headset)
	router.POST("/agent", agent)
	router.StaticFS("/", http.Dir("./public"))
	initDB()
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, req)
}

func report(c *gin.Context) {
	report := findAllHeadsets()
	c.JSON(200, report)
}

func headset(c *gin.Context) {
	var req Headset
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	req.insertHeadset()
	c.JSON(200, req)
}

func agent(c *gin.Context) {
	var req AssignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	h := Headset{Id: req.HeadsetId, Agente: req.Agent}
	h.assingHeadsetToAgent()
	c.JSON(200, h)
}
