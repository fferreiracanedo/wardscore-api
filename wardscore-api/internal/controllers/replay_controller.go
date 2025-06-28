package controllers

import (
	"net/http"
	"strconv"
	"wardscore-api/internal/models"
	"wardscore-api/internal/services"

	"github.com/gin-gonic/gin"
)

// ReplayController gerencia operações relacionadas aos replays
type ReplayController struct {
    replayService *services.ReplayService
}

// NewReplayController cria nova instância do controller
func NewReplayController(replayService *services.ReplayService) *ReplayController {
    return &ReplayController{
        replayService: replayService,
    }
}

// UploadReplay simula upload de replay
// POST /api/v1/replays/upload
func (rc *ReplayController) UploadReplay(c *gin.Context) {
    userID := c.GetHeader("X-User-ID")
    
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "User ID é obrigatório no header X-User-ID",
        })
        return
    }

    id, err := strconv.ParseUint(userID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "User ID inválido",
        })
        return
    }

    var req struct {
        FileName string `json:"file_name" binding:"required"`
        GameID   string `json:"game_id" binding:"required"`
        Champion string `json:"champion"`
        Queue    string `json:"queue"`
        Duration int    `json:"duration"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Dados inválidos: " + err.Error(),
        })
        return
    }

    replay := &models.Replay{
        UserID:       uint(id),
        FileName:     req.FileName,
        OriginalName: req.FileName,
        MatchID:      req.GameID, // GameID vira MatchID
        Champion:     req.Champion,
        Queue:        req.Queue,
        Duration:     req.Duration,
        Status:       models.StatusUploaded,
        FileSize:     1024000, // Mock: 1MB
        FilePath:     "/replays/" + req.FileName,
    }

    createdReplay, err := rc.replayService.Create(replay)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao salvar replay: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "data":    createdReplay,
        "message": "Replay uploaded com sucesso! Será processado em breve.",
    })
}

// GetReplays lista replays do usuário
// GET /api/v1/replays?user_id=1&page=1&limit=10
func (rc *ReplayController) GetReplays(c *gin.Context) {
    userID := c.Query("user_id")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 50 {
        limit = 10
    }

    var replays []models.Replay
    var err error

    if userID != "" {
        id, parseErr := strconv.ParseUint(userID, 10, 32)
        if parseErr != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "success": false,
                "error":   "User ID inválido",
            })
            return
        }
        replays, err = rc.replayService.GetByUserID(uint(id))
    } else {
        // Para listar todos, vamos retornar mensagem de que precisa especificar usuário
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Parâmetro user_id é obrigatório",
        })
        return
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao buscar replays: " + err.Error(),
        })
        return
    }

    // Paginação manual básica
    total := len(replays)
    start := (page - 1) * limit
    end := start + limit
    
    if start >= total {
        replays = []models.Replay{}
    } else {
        if end > total {
            end = total
        }
        replays = replays[start:end]
    }

    totalPages := (total + limit - 1) / limit
    hasNext := page < totalPages
    hasPrev := page > 1

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    replays,
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

// GetReplay busca replay específico
// GET /api/v1/replays/:id
func (rc *ReplayController) GetReplay(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "ID inválido",
        })
        return
    }

    replay, err := rc.replayService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Replay não encontrado",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    replay,
    })
}

// UpdateReplay atualiza dados do replay
// PUT /api/v1/replays/:id
func (rc *ReplayController) UpdateReplay(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "ID inválido",
        })
        return
    }

    var req struct {
        Status     string `json:"status"`
        Champion   string `json:"champion"`
        Queue      string `json:"queue"`
        Duration   int    `json:"duration"`
        MatchDate  string `json:"match_date"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Dados inválidos: " + err.Error(),
        })
        return
    }

    replay, err := rc.replayService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Replay não encontrado",
        })
        return
    }

    if req.Status != "" {
        replay.Status = models.ReplayStatus(req.Status)
    }
    if req.Champion != "" {
        replay.Champion = req.Champion
    }
    if req.Queue != "" {
        replay.Queue = req.Queue
    }
    if req.Duration > 0 {
        replay.Duration = req.Duration
    }

    updatedReplay, err := rc.replayService.Update(replay)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao atualizar replay: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    updatedReplay,
        "message": "Replay atualizado com sucesso",
    })
}

// DeleteReplay remove replay
// DELETE /api/v1/replays/:id
func (rc *ReplayController) DeleteReplay(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "ID inválido",
        })
        return
    }

    err = rc.replayService.Delete(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao deletar replay: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Replay deletado com sucesso",
    })
} 