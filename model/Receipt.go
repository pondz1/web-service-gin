package model

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	db "web-service-gin/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Receipt struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Status      string             `json:"status"`
	LineId      string             `json:"lineid"`
	PicturePath string             `json:"picturePath"`
}

func Upload(c *gin.Context) {

	_, header, err := c.Request.FormFile("upload")
	lineid := c.Request.Form.Get("lineid")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "no file"})
		return
	}
	filename := header.Filename
	err = c.SaveUploadedFile(header, "./image/"+filename)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.RECEIPT)
	collUser := client.Database(db.DB).Collection(db.USER)
	var user User
	err = collUser.FindOne(context.TODO(),
		bson.M{"lineid": lineid}).Decode(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not found user"})
		return
	}
	curser, err := collection.InsertOne(context.TODO(), Receipt{
		ID:          primitive.NewObjectID(),
		Status:      "pending",
		LineId:      lineid,
		PicturePath: "/image/" + filename,
	})
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	id := curser.InsertedID
	c.IndentedJSON(http.StatusOK, gin.H{"message": "uploaded", "id": id})
}

func UpdateStatus(c *gin.Context, client *mongo.Client, receiptId string, lineId string) error {
	var receipt Receipt
	objID, err := primitive.ObjectIDFromHex(receiptId)
	if err != nil {
		return err
	}
	collection := client.Database(db.DB).Collection(db.RECEIPT)
	err = collection.FindOne(context.TODO(),
		bson.M{"_id": objID, "lineid": lineId}).Decode(&receipt)
	if err != nil {
		return err
	}
	fmt.Printf("receipt: %v\n", receipt)
	if receipt.Status == "complete" {
		return errors.New("error: status complete")
	}
	objID2, err := primitive.ObjectIDFromHex(receipt.ID.Hex())
	if err != nil {
		return err
	}
	curser, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID2},
		bson.M{"$set": bson.M{
			"status": "complete",
		}},
	)
	id := curser.UpsertedID
	fmt.Printf("id: %v\n", id)
	if err != nil {
		return err
	}

	return err
}

func GetReceipt(c *gin.Context) {
	var receipt []Receipt

	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.RECEIPT)
	curser, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	curser.All(context.TODO(), &receipt)
	defer curser.Close(context.TODO())
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": receipt})
}

func GetReceiptByUserID(c *gin.Context) {
	var receipt []Receipt
	id := c.Param("id")
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.RECEIPT)
	curser, err := collection.Find(context.TODO(), bson.M{"lineid": id})

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	curser.All(context.TODO(), &receipt)
	defer curser.Close(context.TODO())
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": receipt})
}
