package main

import (
	"log"
	"tulip/backend/controllers"
	"tulip/backend/middlewares"
	"tulip/backend/models"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	authEnforcer, err := casbin.NewEnforcer("./casbin.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	models.ConnectDB()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	public := r.Group("/api")

	// Available from outside
	// /register for convenience, final version should not have it
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	// Available to only logged id + authorization
	userApi := r.Group("/api/user")
	userApi.Use(middlewares.JwtAuthMiddleware())
	userApi.Use(middlewares.Authorizer(authEnforcer))
	userApi.GET("", controllers.GetUsers)
	userApi.GET("/me", controllers.CurrentUser)
	userApi.GET("/:id", controllers.GetUserByID)
	userApi.POST("", controllers.Register)
	userApi.PUT("/:id", controllers.ChangeUser)

	transactionApi := r.Group("api/transaction")
	transactionApi.Use(middlewares.JwtAuthMiddleware())
	transactionApi.Use(middlewares.Authorizer(authEnforcer))
	transactionApi.GET("", controllers.GetTransactions)
	transactionApi.GET("/:id", controllers.GetTransactionById)

	collectionApi := r.Group("/api/collection")
	// TODO: add mw after fix cors authorization
	collectionApi.POST("/add", controllers.AddCollection)
	collectionApi.GET("/getallcol", controllers.GetAllCollect)
	collectionApi.GET("/:name", controllers.GetCollectByName)

	keyApi := r.Group("/api/key")
	keyApi.GET("/all", controllers.GetAllKeys)
	keyApi.POST("/add", controllers.AddKey)
	keyApi.DELETE("/remove", controllers.RemoveKey)
	keyApi.PUT("/update", controllers.UpdateKey)

	nftsApi := r.Group("/api/nfts")
	// nftsApi.Use(middlewares.JwtAuthMiddleware())
	// nftsApi.Use(middlewares.Authorizer(authEnforcer))
	nftsApi.GET("", controllers.GetNFTs)
	nftsApi.GET("/:id", controllers.GetNFTById)
	nftsApi.GET("/shownft/:name", controllers.GetNFTsByCollection)
	nftsApi.GET("/minthistory", controllers.GetNFTMintHistory)
	nftsApi.POST("/mintfilenfts", controllers.MintFileNFT)
	nftsApi.POST("/minttextnfts", controllers.MintTextNFT)
	// TODO: POST /

	manualTransferApi := r.Group("/api/manualtransfer")
	manualTransferApi.Use(middlewares.JwtAuthMiddleware())
	manualTransferApi.Use(middlewares.Authorizer(authEnforcer))
	manualTransferApi.GET("", controllers.GetManualTransfers)
	// TODO: do this with VSYS SDK
	//manualTransferApi.POST("", controllers.ManualTransfer)
	manualTransferApi.GET("/:id", controllers.GetManualTransferById)

	contractApi := r.Group("/api/contract")
	contractApi.Use(middlewares.JwtAuthMiddleware())
	contractApi.Use(middlewares.Authorizer(authEnforcer))
	contractApi.GET("", controllers.GetContracts)
	contractApi.GET("/:id", controllers.GetContractById)
	contractApi.POST("", controllers.AddContract)
	contractApi.PUT("/:id", controllers.UpdateContract)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error on Server Launch", err.Error())
	}
}
