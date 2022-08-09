package controllers

import (
	"fmt"
	"net/http"
	"tulip/backend/models"

	"github.com/gin-gonic/gin"
)

type CollectionInput struct {
	CollectionName string `json:"collection_name" binding:"required"`
	MetaInfoType   string `json:"metainfo_type" binding:"required"`
	DesFileKey     string `json:"des_file_key" binding:"required"`
}

func AddCollection(c *gin.Context) {
	var input CollectionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := models.Collection{}

	m.CollectionName = input.CollectionName
	m.MetaInfoType = input.MetaInfoType
	m.DesFileKey = input.DesFileKey

	_, err := m.InsertCollect()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("here")

	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
}

func GetAllCollect(c *gin.Context) {
	collections, err := models.GetAllCollect()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(collections)
	c.JSON(http.StatusOK, collections)
}

func GetCollectByName(c *gin.Context) {
	name, berr := c.Params.Get("name")
	if !berr {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "params wrong"})
		// c.Abort()
		return
	}
	collection, err := models.GetCollectByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
		// c.Abort()
		return
	}
	metaTemplate, err := models.GetTemplateByName(collection.MetaInfoType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "template not found"})
		// c.Abort()
		return
	}
	ret := gin.H{
		"collection_info":    collection,
		"meta_template_info": metaTemplate,
	}
	c.JSON(http.StatusOK, ret)
}
