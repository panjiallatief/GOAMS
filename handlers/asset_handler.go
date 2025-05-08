package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/service"
)

// AssetHandler handles HTTP requests for assets
type AssetHandler struct {
	service service.AssetService
}

// NewAssetHandler registers asset routes on the router
func NewAssetHandler(r *gin.Engine, s service.AssetService) {
	h := &AssetHandler{service: s}
	assets := r.Group("/assets")
	{
		assets.POST("", h.CreateAsset)
		assets.GET("", h.GetAssets)
		assets.GET(":id", h.GetAsset)
		assets.PUT(":id", h.UpdateAsset)
		assets.DELETE(":id", h.DeleteAsset)
	}
}

// CreateAsset handles creating a new asset
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateAsset(&asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, asset)
}

// GetAssets returns all assets
func (h *AssetHandler) GetAssets(c *gin.Context) {
	assets, err := h.service.GetAllAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assets)
}

// GetAsset returns a single asset by ID
func (h *AssetHandler) GetAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}
	asset, err := h.service.GetAssetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}
	c.JSON(http.StatusOK, asset)
}

// UpdateAsset handles updating an existing asset
func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	asset.IDAsset = uint(id)
	if err := h.service.UpdateAsset(&asset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, asset)
}

// DeleteAsset handles deleting an asset
func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}
	if err := h.service.DeleteAsset(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
