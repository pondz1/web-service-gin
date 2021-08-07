package model

import (
	"context"
	"net/http"
	db "web-service-gin/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty" swaggerignore:"true"`
	Point   float64            `json:"p_point"`
	ProName string             `json:"p_name"`
}

type Response struct {
	Code   int    `json:code`   // Response Code
	Status string `json:status` // Response Status
}

type Example struct {
	Id      string `json:id`      // Response Code
	Message string `json:message` // Response Status
}

// InsertProduct godoc
// @summary Insert Product
// @description Insert product
// @security ApiKeyAuth
// @id InsertProduct
// @accept json
// @produce json
// @param Product body Product true "Product data to be Insert"
// @response 201 {object} Example "OK"
// @response 400 {object} Example "Bad Request"
// @Router /insertProduct [post]
func InsertProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if len(newProduct.ProName) == 0 || newProduct.Point < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name of product is null or point must be greater than 0"})
		return
	}
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.PRODUCT)
	curser, err := collection.InsertOne(context.TODO(), Product{
		ID:      primitive.NewObjectID(),
		Point:   newProduct.Point,
		ProName: newProduct.ProName})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := curser.InsertedID

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "inserted!", "id": id})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product Product
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.PRODUCT)
	err2 := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)

	if err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "found", "data": product})

}

func GetAllProducts(c *gin.Context) {
	var products []Product

	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.PRODUCT)
	curser, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	curser.All(context.TODO(), &products)
	// for curser.Next(context.TODO()) {
	// 	var u User
	// 	curser.Decode(&u)
	// 	users = append(users, u)
	// }
	defer curser.Close(context.TODO())
	c.IndentedJSON(http.StatusOK, gin.H{"data": products})

}
