package data

import (
	"github.com/gin-gonic/gin"
	"tod/apikey"
)

type handler struct {
	repo    Repository
	repoApi apikey.Repository
}

func NewHandler(repo Repository, repoApi apikey.Repository) *handler {
	return &handler{repo, repoApi}
}

func (h *handler) CreateDataHandler(c *gin.Context) {
	qryApi := c.Query("apikey")
	if qryApi == "" {
		c.JSON(400, gin.H{
			"message": "apikey is required",
		})
		return
	}

	checkApi, err := h.repoApi.CheckApikey(c, qryApi)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !checkApi {
		c.JSON(400, gin.H{
			"message": "apikey is invalid",
		})
		return
	}

	var input DataInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	ctx := c.Request.Context()
	data, err := h.repo.CreateData(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}

func (h *handler) GetDataHandler(c *gin.Context) {
	qryApi := c.Query("apikey")
	if qryApi == "" {
		c.JSON(400, gin.H{
			"message": "apikey is required",
		})
		return
	}

	checkApi, err := h.repoApi.CheckApikey(c, qryApi)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !checkApi {
		c.JSON(400, gin.H{
			"message": "apikey is invalid",
		})
		return
	}

	ctx := c.Request.Context()
	typeName := c.Query("type")
	switch typeName {
	case "truth":
		data, err := h.repo.GetAllTruthData(ctx)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": data,
		})
	case "dare":
		data, err := h.repo.GetAllDareData(ctx)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"data": data,
		})
	default:
		c.JSON(400, gin.H{
			"message": "type is invalid",
		})
	}
}