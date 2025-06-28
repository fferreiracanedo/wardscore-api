package controllers

import (
	"net/http"
	"strconv"
	"wardscore-api/internal/models"
	"wardscore-api/internal/services"

	"github.com/gin-gonic/gin"
)

// UserController gerencia operações relacionadas aos usuários
type UserController struct {
    userService *services.UserService
}

// NewUserController cria nova instância do controller
func NewUserController(userService *services.UserService) *UserController {
    return &UserController{
        userService: userService,
    }
}

// CreateUser cria um novo usuário
// POST /api/v1/users
func (uc *UserController) CreateUser(c *gin.Context) {
    var req struct {
        RiotID   string `json:"riot_id" binding:"required"`
        GameName string `json:"game_name" binding:"required"`
        TagLine  string `json:"tag_line" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Region   string `json:"region"`
        PUUID    string `json:"puuid"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Dados inválidos: " + err.Error(),
        })
        return
    }

    user := &models.User{
        RiotID:   req.RiotID,
        GameName: req.GameName,
        TagLine:  req.TagLine,
        Email:    req.Email,
        Region:   req.Region,
        PUUID:    req.PUUID,
    }

    if user.Region == "" {
        user.Region = "BR1"
    }

    createdUser, err := uc.userService.Create(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao criar usuário: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "data":    createdUser,
        "message": "Usuário criado com sucesso",
    })
}

// GetAllUsers lista todos os usuários com paginação
// GET /api/v1/users?page=1&limit=10
func (uc *UserController) GetAllUsers(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 100 {
        limit = 10
    }

    users, total, err := uc.userService.GetAll(page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao buscar usuários: " + err.Error(),
        })
        return
    }

    totalPages := (total + int64(limit) - 1) / int64(limit)
    hasNext := page < int(totalPages)
    hasPrev := page > 1

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    users,
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

// GetProfile busca perfil do usuário
// GET /api/v1/users/profile
func (uc *UserController) GetProfile(c *gin.Context) {
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

    user, err := uc.userService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Usuário não encontrado",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    user,
    })
}

// UpdateProfile atualiza perfil do usuário
// PUT /api/v1/users/profile
func (uc *UserController) UpdateProfile(c *gin.Context) {
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
        GameName  string `json:"game_name"`
        TagLine   string `json:"tag_line"`
        Email     string `json:"email"`
        AvatarURL string `json:"avatar_url"`
        Region    string `json:"region"`
        PUUID     string `json:"puuid"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Dados inválidos: " + err.Error(),
        })
        return
    }

    user, err := uc.userService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Usuário não encontrado",
        })
        return
    }

    if req.GameName != "" {
        user.GameName = req.GameName
    }
    if req.TagLine != "" {
        user.TagLine = req.TagLine
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    if req.AvatarURL != "" {
        user.AvatarURL = req.AvatarURL
    }
    if req.Region != "" {
        user.Region = req.Region
    }
    if req.PUUID != "" {
        user.PUUID = req.PUUID
    }

    updatedUser, err := uc.userService.Update(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao atualizar usuário: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    updatedUser,
        "message": "Perfil atualizado com sucesso",
    })
}

// DeleteUser remove usuário
// DELETE /api/v1/users/:id
func (uc *UserController) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "ID inválido",
        })
        return
    }

    err = uc.userService.Delete(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Falha ao deletar usuário: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Usuário deletado com sucesso",
    })
} 