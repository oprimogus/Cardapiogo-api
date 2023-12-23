package controller

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"github.com/oprimogus/cardapiogo/internal/domain/user"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	log = logger.GetLogger("UserController")
	validate = validator.New()
)

type userController struct {
	service user.Service
}

func NewUserController(repository user.Repository) *userController {
	return &userController{service: user.NewService(repository)}
}

func (c *userController) CreateUserHandler(ctx *gin.Context) {
	var userParams user.CreateUserParams
	err :=  ctx.BindJSON(&userParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	// Validação
	validationErr := validate.Struct(userParams)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	userCreated := c.service.CreateUser(ctx, userParams)
	if userCreated != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": userCreated})
	}
	ctx.Status(http.StatusCreated)

	// Resto do seu código para criar o usuário...
}
