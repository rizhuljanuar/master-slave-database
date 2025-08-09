package routes

import (
	
	"net/http"

	"github.com/gin-gonic/gin"

	"apps/models"
	"apps/handlers"
)

func SetupRoutes(r *gin.Engine, h *handlers.ItemHandler) {
	// Handler untuk rute yang tidak ditemukan
	r.NoRoute(func(c *gin.Context) {
		rsp := &models.Response{}
		rsp.WithCode(http.StatusNotFound)
		rsp.WithMessage(`Rute tidak ditemukan`)
		c.JSON(http.StatusNotFound, rsp)
	})

	r.POST("/items", func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			rsp := &models.Response{}
			rsp.WithCode(http.StatusBadRequest)
			rsp.WithMessage(`Data tidak valid`)
			c.JSON(rsp.Code, rsp)
			return
		}

		rsp := h.Store(item)

		c.JSON(rsp.Code, rsp)
	})

	r.GET("/items", func(c *gin.Context) {
	
		rsp := h.FindAll()

		c.JSON(rsp.Code, rsp)
	})

	r.GET("/items/:id", func(c *gin.Context) {
		id := c.Param("id")

		rsp := h.FindByID(id)

		c.JSON(rsp.Code, rsp)
	})

	r.PUT("/items/:id", func(c *gin.Context) {
		id := c.Param("id")
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			rsp := &models.Response{}
			rsp.WithCode(http.StatusBadRequest)
			rsp.WithMessage(`Data tidak valid`)
			c.JSON(rsp.Code, rsp)
			return
		}

		rsp := h.Update(id, item)

		c.JSON(rsp.Code, rsp)
	})

	r.DELETE("/items/:id", func(c *gin.Context) {
		id := c.Param("id")
		rsp := h.Delete(id)

		c.JSON(rsp.Code, rsp)
	})
}