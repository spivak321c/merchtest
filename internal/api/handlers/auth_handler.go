package handlers

import (
	"encoding/json"
	"net/http"

	"api-customer-merchant/internal/db/models"
	"api-customer-merchant/internal/domain/merchant"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service *merchant.MerchantService
}

func NewMerchantAuthHandler(s *merchant.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: s}
}

// POST /merchant/apply
// Used by a prospective merchant to submit an application.
func (h *MerchantHandler) Apply(c *gin.Context) {
	var req struct {
		models.MerchantBasicInfo
		PersonalAddress            map[string]any `json:"personal_address" validate:"required"`
		WorkAddress                map[string]any `json:"work_address" validate:"required"`
		models.MerchantBusinessInfo
		models.MerchantDocuments
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// Convert personal_address and work_address to JSONB
	personalAddressJSON, err := json.Marshal(req.PersonalAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid personal_address JSON: " + err.Error()})
		return
	}
	workAddressJSON, err := json.Marshal(req.WorkAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid work_address JSON: " + err.Error()})
		return
	}

	app := &models.MerchantApplication{
		MerchantBasicInfo:    req.MerchantBasicInfo,
		MerchantAddress:      models.MerchantAddress{PersonalAddress: personalAddressJSON, WorkAddress: workAddressJSON},
		MerchantBusinessInfo: req.MerchantBusinessInfo,
		MerchantDocuments:    req.MerchantDocuments,
	}

	app, err = h.service.SubmitApplication(c.Request.Context(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit application: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, app)
}

// GET /merchant/application/:id
// Allows an applicant to view their application status.
func (h *MerchantHandler) GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := h.service.GetApplication(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "application not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

// GET /merchant/me
// Fetches the current user's merchant account (if approved).
func (h *MerchantHandler) GetMyMerchant(c *gin.Context) {
	userID, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	m, err := h.service.GetMerchantByUserID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}