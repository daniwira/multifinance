package handler

import (
	"net/http"
	"strconv"

	"github.com/daniwira/multifinance/internal/domain/limit"
	"github.com/daniwira/multifinance/internal/service"
	"github.com/gin-gonic/gin"
)

type LimitHandler struct {
	limitService service.LimitService
}

func NewLimitHandler(limitService service.LimitService) *LimitHandler {
	return &LimitHandler{
		limitService: limitService,
	}
}

func (h *LimitHandler) GetLimits(c *gin.Context) {
	limits, err := h.limitService.GetLimits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, limits)
}

func (h *LimitHandler) GetLimit(c *gin.Context) {
	id := c.Param("id")

	limit, err := h.limitService.GetLimit(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Limit not found"})
		return
	}

	c.JSON(http.StatusOK, limit)
}

func (h *LimitHandler) CreateLimit(c *gin.Context) {
	var limit limit.Limit
	err := c.BindJSON(&limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLimit, err := h.limitService.CreateLimit(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdLimit)
}

func (h *LimitHandler) UpdateLimit(c *gin.Context) {
	id := c.Param("id")

	var updatedLimit limit.Limit
	err := c.BindJSON(&updatedLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limitID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLimit.ID = limitID

	limit, err := h.limitService.UpdateLimit(updatedLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, limit)
}

func (h *LimitHandler) DeleteLimit(c *gin.Context) {
	id := c.Param("id")

	err := h.limitService.DeleteLimit(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
