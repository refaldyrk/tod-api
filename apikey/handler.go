package apikey

import "github.com/gin-gonic/gin"

type handler struct {
	repo Repository
}

func NewHandler(repo Repository) *handler {
	return &handler{repo: repo}
}

func (h *handler) GetAllApikeyHandler(c *gin.Context) {
	apikeyyyy := c.Query("apikey")
	if apikeyyyy != "root" {
		c.JSON(400, gin.H{
			"error": "Invalid Apikey",
		})
		return
	}

	models, err := h.repo.GetAllApikey(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": models})
}

func (h *handler) CreateApikeyHandler(c *gin.Context) {
	apikeyyyy := c.Query("apikey")
	if apikeyyyy != "root" {
		c.JSON(400, gin.H{
			"error": "Invalid Apikey",
		})
		return
	}

	var input ApikeyInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	model, err := h.repo.CreateApikey(c, input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": model,
	})

	return
}

func (h *handler) DeleteApikeyHandler(c *gin.Context) {
	apikeyyyy := c.Query("apikey")
	if apikeyyyy != "root" {
		c.JSON(400, gin.H{
			"error": "Invalid Apikey",
		})
		return
	}

	key := c.Query("key")

	err := h.repo.DeleteApikey(c, key)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "success",
	})

	return
}

func (h *handler) CheckApikeyHandler(c *gin.Context) {
	apikeyyyy := c.Query("apikey")
	if apikeyyyy != "root" {
		c.JSON(400, gin.H{
			"error": "Invalid Apikey",
		})
		return
	}

	key := c.Query("key")

	ok, err := h.repo.CheckApikey(c, key)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ok,
	})

	return
}
