package handler

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/BENODR27/url-shortener-go/internal/service"
)

type URLHandler struct {
    service service.URLService
}

func NewURLHandler(service service.URLService) *URLHandler {
    return &URLHandler{service: service}
}

func (h *URLHandler) Shorten(c *gin.Context) {
    var req struct{ URL string `json:"url" binding:"required,url"` }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    code, err := h.service.Shorten(context.Background(), req.URL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"short_url": c.Request.Host + "/" + code})
}

func (h *URLHandler) Resolve(c *gin.Context) {
    code := c.Param("code")
    original, err := h.service.Resolve(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }
    c.Redirect(http.StatusFound, original)
}
