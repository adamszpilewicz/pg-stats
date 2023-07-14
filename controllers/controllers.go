package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stats-pg/repository"
)

type Controller struct {
	DbHandler repository.Repository
}

func (c *Controller) GetStats(ctx *gin.Context) {
	stats, err := c.DbHandler.GetStats()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, stats)
}

func (c *Controller) GetStatsByName(ctx *gin.Context) {
	dbName := ctx.Param("dbName")
	stats, err := c.DbHandler.GetStatsByName(ctx, dbName)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, stats)
}

func NewController(dbHandler repository.Repository) *Controller {
	return &Controller{
		DbHandler: dbHandler,
	}
}
