package controllers

import (
	"linkbaby/models"
	"linkbaby/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.Usuario

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(newUser)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "email already in use"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email string `json:"email" binding:"required"`
		Senha string `json:"senha" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.LoginUser(loginData.Email, loginData.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user.Senha = ""
	c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c *gin.Context) {
	email := c.DefaultQuery("email", "")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	var user models.Usuario

	collection := models.DB.Database("linkbaby").Collection("usuarios")
	err := collection.FindOne(models.Ctx, map[string]interface{}{"email": email}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	user.Senha = ""

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	deleted, _ := services.DeleteUser(id)
	if deleted {
		c.JSON(http.StatusOK, gin.H{"message": "Usuário removido com sucesso"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
	}
}
