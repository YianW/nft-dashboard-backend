package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"tulip/backend/models"
	"tulip/backend/utils/token"
)

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	u, err := models.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, u)
}

type ChangeUserInput struct {
	Username    string `json:"username"`
	OldPassword string `json:"old-password"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

func ChangeUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var input ChangeUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	tokenUID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	u, err := models.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if tokenUID != u.ID && u.Role != "superadmin" {
		// TODO: do more careful access control, for now you can change your info and superadmin can change anyone
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you can change your own data only"})
		c.Abort()
		return
	}
	if input.Username != "" {
		u.Username = input.Username
	}
	if input.Password != "" {
		if input.OldPassword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "When password changed 'old-password' needed"})
			c.Abort()
			return
		}
		if err := models.VerifyPassword(u.Password, input.OldPassword); err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				c.JSON(http.StatusBadRequest, gin.H{"error": "old password incorrect"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}
		// TODO: Add password verification or move to models
		u.Password = input.Password
	}
	if input.Role != "" {
		// TODO: additional verification of role change. Only superadmin can make admin and so on
		u.Role = input.Role
	}
	newU, err := u.UpdateUser()
	newU.PrepareGive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newU)
}
