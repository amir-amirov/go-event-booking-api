package routes

import (
	"net/http"

	"github.com/amir-amirov/go-event-booking-api/models"
	"github.com/amir-amirov/go-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	var err error

	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func loginUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	jwt, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": jwt})

}
