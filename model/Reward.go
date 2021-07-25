package model

import (
	"context"
	"fmt"
	"net/http"
	db "web-service-gin/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reward struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	LineId      string             `json:"r_lineid"`
	ItemProduct Product            `json:"r_product"`
}

func InsertReward(c *gin.Context) {
	var newReward Reward
	var product Product
	var user User
	if err := c.BindJSON(&newReward); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.REWARD)
	collProduct := client.Database(db.DB).Collection(db.PRODUCT)
	collUser := client.Database(db.DB).Collection(db.USER)
	objID, _ := primitive.ObjectIDFromHex(newReward.ItemProduct.ID.Hex())
	err = collProduct.FindOne(context.TODO(),
		bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found product"})
		return
	}
	err = collUser.FindOne(context.TODO(),
		bson.M{"lineid": newReward.LineId}).Decode(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found user"})
		return
	}
	point := user.Points - product.Point
	if user.Points < product.Point {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not enough points"})
		return
	}
	curser, err := collection.InsertOne(context.TODO(), Reward{
		ID:          primitive.NewObjectID(),
		LineId:      newReward.LineId,
		ItemProduct: newReward.ItemProduct,
	})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = collUser.FindOneAndUpdate(context.TODO(),
		bson.M{"lineid": newReward.LineId}, bson.M{"$set": bson.M{"points": point}}).Decode(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := curser.InsertedID
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "inserted!", "id": id})
}

func GetRewardByUserId(c *gin.Context) {
	id := c.Param("id")
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.REWARD)
	curser, err := collection.Find(context.TODO(), bson.M{"lineid": id})
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "no data"})
		return
	}
	var rewards []bson.M
	for curser.Next(context.TODO()) {
		var reward Reward
		fmt.Println(curser.Decode(&reward))
		rewards = append(rewards, bson.M{
			"productName": reward.ItemProduct.ProName,
			"usePoint":    reward.ItemProduct.Point,
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": rewards})
}
