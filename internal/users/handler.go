package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/natanfds/epic-dice/utils"
)

type UserHandler struct {
	repository *UserRepository
}

func (u *UserHandler) Create(g *gin.Context) {
	var createData CreateUserDTO
	bodyData := g.ShouldBind(&createData)
	validationErr := utils.Validate.Struct(createData)
	if bodyData != nil || validationErr != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hashedPassword, err := utils.Password.Encrypt(createData.Password)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	if err := u.repository.Create(UserModel{
		Username:    createData.Username,
		DisplayName: createData.DisplayName,
		Email:       createData.Email,
		Hash:        hashedPassword,
		Color:       createData.Color,
	}); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (u *UserHandler) Update(g *gin.Context) {

}

func (u *UserHandler) Delete(g *gin.Context) {

}

func (u *UserHandler) Login(g *gin.Context) {

}

func NewUserHandler(repository *UserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}
