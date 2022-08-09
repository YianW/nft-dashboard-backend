package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tulip/backend/models"
)

func GetContracts(c *gin.Context) {
	contracts, err := models.GetContracts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func GetContractById(c *gin.Context) {
	id := c.Param("id")
	contract, err := models.GetContractById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, contract)
}

type AddContractInput struct {
	ContractId                string `json:"contract_id" binding:"required"`
	Name                      string `json:"name" binding:"required"`
	RegistrationTransactionId string `json:"registration_transaction_id" binding:"required"`
	Maker                     string `json:"maker" binding:"required"`
	ContractDescription       string `json:"contract_description"`
	Active                    bool   `json:"active"`
}

func AddContract(c *gin.Context) {
	var input AddContractInput
	// Default value
	input.Active = true

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrt := models.Contract{}
	ctrt.ContractId = input.ContractId
	ctrt.Name = input.Name
	ctrt.RegistrationTransactionId = input.RegistrationTransactionId
	ctrt.Maker = input.Maker
	ctrt.ContractDescription = input.ContractDescription
	ctrt.Active = input.Active

	_, err := ctrt.SaveContract()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}

type ChangeContractInput struct {
	Name                string `json:"name"`
	ContractDescription string `json:"contract_description"`
	Active              bool   `json:"active"`
}

func UpdateContract(c *gin.Context) {
	id := c.Param("id")

	var input ChangeContractInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrt, err := models.GetContractById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// TODO: Formally define what and how can be changed
	if input.Name != "" {
		ctrt.Name = input.Name
	}
	if input.ContractDescription != "" {
		ctrt.ContractDescription = input.ContractDescription
	}
	if input.Active != ctrt.Active {
		ctrt.Active = input.Active
	}
	new_ctrt, err := ctrt.UpdateContract()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, new_ctrt)
}
