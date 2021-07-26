package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	m "web-service-gin/model"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"massage": "Successfully connected and pinged."})
	})
	router.Static("/image", "./image")
	router.GET("/getuser", m.GetUsers)
	router.GET("/userid/:id", m.GetUsersByID)
	router.GET("/getAllProducts", m.GetAllProducts)
	router.GET("/getProduct/:id", m.GetProductByID)
	router.GET("/getRewardByUserId/:id", m.GetRewardByUserId)
	router.GET("/getReceipt", m.GetReceipt)
	router.GET("/getReceipt/:id", m.GetReceiptByUserID)
	router.POST("/insertUser", m.InsertUser)
	router.POST("/updateUser", m.UpdateUser)
	router.POST("/addUserPoint/:receipt", m.UpdateUserPoint)
	router.POST("/insertProduct", m.InsertProduct)
	router.POST("/insertReward", m.InsertReward)
	router.POST("/uploadReceipt", m.Upload)

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
