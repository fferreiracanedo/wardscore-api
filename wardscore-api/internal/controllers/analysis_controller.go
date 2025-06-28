package controllers

import (
	"net/http"
	"strconv"
	"wardscore-api/internal/models"
	"wardscore-api/internal/services"

	"github.com/gin-gonic/gin"
)

// AnalysisController gerencia operações relacionadas às análises
type AnalysisController struct {
    analysisService *services.AnalysisService
}

// NewAnalysisController cria nova instância do controller
func NewAnalysisController(analysisService *services.AnalysisService) *AnalysisController {
    return &AnalysisController{
        analysisService: analysisService,
    }
}

// GetAnalysis busca análise por ID
// GET /api/v1/analysis/:id
func (ac *AnalysisController) GetAnalysis(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "ID inválido",
        })
        return
    }

    analysis, err := ac.analysisService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Análise não encontrada",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    analysis,
    })
}

// ProcessReplay processa um replay e gera análise
// POST /api/v1/analysis/process/:replay_id
func (ac *AnalysisController) ProcessReplay(c *gin.Context) {
    replayIDParam := c.Param("replay_id")
    replayID, err := strconv.ParseUint(replayIDParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Replay ID inválido",
        })
        return
    }

    userID := c.GetHeader("X-User-ID")
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "User ID é obrigatório no header X-User-ID",
        })
        return
    }

    _, err = strconv.ParseUint(userID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "User ID inválido",
        })
        return
    }

    // Processar replay e criar análise
    analysis, err := ac.analysisService.ProcessReplay(uint(replayID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao processar replay: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "data":    analysis,
        "message": "Replay processado com sucesso! Análise gerada.",
    })
}

// GetUserAnalyses busca todas as análises de um usuário
// GET /api/v1/analysis/user/:user_id?page=1&limit=10
func (ac *AnalysisController) GetUserAnalyses(c *gin.Context) {
    userIDParam := c.Param("user_id")
    userID, err := strconv.ParseUint(userIDParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "User ID inválido",
        })
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 50 {
        limit = 10
    }

    analyses, err := ac.analysisService.GetByUserID(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao buscar análises: " + err.Error(),
        })
        return
    }

    // Paginação manual básica
    total := len(analyses)
    start := (page - 1) * limit
    end := start + limit
    
    if start >= total {
        analyses = []models.Analysis{}
    } else {
        if end > total {
            end = total
        }
        analyses = analyses[start:end]
    }

    totalPages := (total + limit - 1) / limit
    hasNext := page < totalPages
    hasPrev := page > 1

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    analyses,
        "meta": gin.H{
            "page":        page,
            "limit":       limit,
            "total":       total,
            "total_pages": totalPages,
            "has_next":    hasNext,
            "has_prev":    hasPrev,
        },
    })
} 