package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	m "web-service-gin/model"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"web-service-gin/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	docs.SwaggerInfo.Version = "1.0"
	router := gin.Default()

	router.Use(cors.Default())

	router.Static("/image", "./image")
	v1 := router.Group("/api/v1")
	{
		v1.GET("ping", ping)
		v1.GET("getuser", m.GetUsers)
		v1.GET("userid/:id", m.GetUsersByID)
		v1.GET("getAllProducts", m.GetAllProducts)
		v1.GET("getProduct/:id", m.GetProductByID)
		v1.GET("getRewardByUserId/:id", m.GetRewardByUserId)
		v1.GET("getReceipt", m.GetReceipt)
		v1.GET("getReceipt/:id", m.GetReceiptByUserID)
		v1.POST("insertUser", m.InsertUser)
		v1.POST("updateUser", m.UpdateUser)
		v1.POST("addUserPoint/:receipt", m.UpdateUserPoint)
		v1.POST("insertProduct", m.InsertProduct)
		v1.POST("insertReward", m.InsertReward)
		v1.POST("uploadReceipt", m.Upload)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(getPort())

}

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}

// ping godoc
// @summary Health Check
// @description Health checking for the service
// @id ping
// @produce json
// @response 200 {object} string "OK"
// @router /ping [get]
func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"massage": "Successfully connected and pinged."})
}
