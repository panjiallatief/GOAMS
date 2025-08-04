package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panjiallatief/GOAMS/service"
)

// DropdownHandler handles requests for dropdown data
type DropdownHandler struct {
	service service.DropdownService
}

// NewDropdownHandler registers dropdown routes on the router
func NewDropdownHandler(r *gin.Engine, s service.DropdownService) {
	h := &DropdownHandler{service: s}

	dropdowns := r.Group("/dropdowns")
	{
		dropdowns.GET("/pics", h.GetPICs)
		dropdowns.GET("/divisions", h.GetDivisions)
		dropdowns.GET("/owners", h.GetOwners)
		dropdowns.GET("/vendors", h.GetVendors)
		dropdowns.GET("/categories", h.GetCategories)
		dropdowns.GET("/brands", h.GetBrands)
		dropdowns.GET("/models", h.GetModels)
	}
}

func (h *DropdownHandler) GetPICs(c *gin.Context) {
	pics, err := h.service.GetAllPICs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pics)
}

func (h *DropdownHandler) GetDivisions(c *gin.Context) {
	divisions, err := h.service.GetAllDivisions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, divisions)
}

func (h *DropdownHandler) GetOwners(c *gin.Context) {
	owners, err := h.service.GetAllOwners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, owners)
}

func (h *DropdownHandler) GetVendors(c *gin.Context) {
	vendors, err := h.service.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vendors)
}

func (h *DropdownHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *DropdownHandler) GetBrands(c *gin.Context) {
	brands, err := h.service.GetAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *DropdownHandler) GetModels(c *gin.Context) {
	models, err := h.service.GetAllModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}
