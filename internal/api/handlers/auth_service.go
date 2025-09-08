package handlers

import (
    "net/http"

    "api-customer-merchant/internal/db/models"
    "api-customer-merchant/internal/domain/merchant"

    "github.com/gin-gonic/gin"
)

type MerchantHandler struct {
    service *merchant.MerchantService
}

func NewMerchantHandler(s *merchant.MerchantService) *MerchantHandler {
    return &MerchantHandler{service: s}
}

// POST /merchant/apply
// Used by a user to submit a merchant application.
func (h *MerchantHandler) Apply(c *gin.Context) {
    var req models.MerchantApplication
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
    }
    app, err := h.service.SubmitApplication(c.Request.Context(), &req)
    if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusCreated, app)
}

// GET /merchant/application/:id
// Allows an applicant to view their application status.
func (h *MerchantHandler) GetApplication(c *gin.Context) {
    id := c.Param("id")
    app, err := h.service.GetApplication(c.Request.Context(), id)
    if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": "application not found"}); return }
    c.JSON(http.StatusOK, app)
}

// GET /merchant/me
// Fetches the current user's merchant account (if approved).
func (h *MerchantHandler) GetMyMerchant(c *gin.Context) {
    userID, ok := c.Get("id")
    if !ok { c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"}); return }
    m, err := h.service.GetMerchantByUserID(c.Request.Context(), userID.(string))
    if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found"}); return }
    c.JSON(http.StatusOK, m)
}