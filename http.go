package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createKv(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "not implemented")
}

func getKv(c *gin.Context) {
	c.Params.ByName("id")

	c.IndentedJSON(http.StatusCreated, c.Params.ByName("id"))
}

func deleteKv(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "not implemented")
}

func updateKv(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "newAlbum")
}
