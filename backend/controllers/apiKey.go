package controllers

import (
	"fmt"
	"net/http"
	"tulip/backend/models"

	"github.com/gin-gonic/gin"
)

type KeyInput struct {
	Key string `json:"key" binding:"required"`
}

func AddKey(c *gin.Context) {
	var input KeyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		return
	}

	m := models.APIkey{}

	m.Key = input.Key

	_, err := m.InsertKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
}

func UpdateKey(c *gin.Context) {
	var input KeyInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		return
	}

	m := models.APIkey{}

	m.Key = input.Key

	_, err := m.UpdateKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
}

type deleteInput struct {
	ID  uint   `json:"ID" binding:"required"`
	Key string `json:"key" binding:"required"`
}

func RemoveKey(c *gin.Context) {
	var input deleteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("111")
		fmt.Println(err.Error())
		return
	}

	m := models.APIkey{}

	m.ID = input.ID
	m.Key = input.Key
	fmt.Println(m.ID)
	fmt.Println(m.Key)

	_, err := m.RemoveKey()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
}

func GetAllKeys(c *gin.Context) {
	keys, err := models.GetAllKeys()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, keys)
}
