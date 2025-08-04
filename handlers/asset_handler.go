package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panjiallatief/GOAMS/models"
	"github.com/panjiallatief/GOAMS/service"
)

// AssetHandler handles HTTP requests for assets
type AssetHandler struct {
	service    service.AssetService
	logService service.LogService
}

// NewAssetHandler registers asset routes on the router
func NewAssetHandler(r *gin.Engine, s service.AssetService, ls service.LogService) {
	h := &AssetHandler{service: s, logService: ls}
	assets := r.Group("/assets")
	{
		assets.POST("", h.CreateAsset)
		assets.GET("", h.GetAssets)
		assets.GET(":id", h.GetAsset)
		assets.PUT(":id", h.UpdateAsset)
		assets.DELETE(":id", h.DeleteAsset)
		assets.GET(":id/logs", h.GetAssetLogs)
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
		if errors.Is(err, service.ErrSNAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Serial Number sudah digunakan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the creation
	h.logService.CreateLog(uint(asset.IDAsset), "Create", "System", fmt.Sprintf("Created asset: %s", asset.ItemName))

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
	asset.IDAsset = int64(id)
	if err := h.service.UpdateAsset(&asset); err != nil {
		if errors.Is(err, service.ErrSNAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Create log entry for asset update
	h.logService.CreateLog(uint(id), "Updated", "System", "Asset information updated")

	c.JSON(http.StatusOK, asset)
}

// DeleteAsset handles deleting an asset
func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Create log entry before deletion
	h.logService.CreateLog(uint(id), "Deleted", "System", "Asset removed from system")

	if err := h.service.DeleteAsset(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetAssetLogs returns activity logs for a specific asset
func (h *AssetHandler) GetAssetLogs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Get real logs from database
	logs, err := h.logService.GetAssetLogs(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}
