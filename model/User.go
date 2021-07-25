package model

import (
	"context"
	"net/http"
	db "web-service-gin/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	LineId     string             `json:"lineid"`
	Name       string             `json:"displayName"`
	Points     float64            `json:"points"`
	PictureUrl string             `json:"pictureUrl"`
}

func GetUsers(c *gin.Context) {
	var users []User

	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.USER)
	curser, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	curser.All(context.TODO(), &users)
	// for curser.Next(context.TODO()) {
	// 	var u User
	// 	curser.Decode(&u)
	// 	fmt.Printf("u._id: %v\n", u.ID)
	// 	users = append(users, u)
	// }
	defer curser.Close(context.TODO())
	c.IndentedJSON(http.StatusOK, gin.H{"data": users})

}
func GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.USER)
	err2 := collection.FindOne(context.TODO(), bson.M{"lineid": id}).Decode(&user)

	if err2 != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "found", "data": user})

}

func InsertUser(c *gin.Context) {
	var newUser User
	var newUser2 User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.USER)
	err2 := collection.FindOne(context.TODO(), bson.M{"lineid": newUser.LineId}).Decode(&newUser2)

	if err2 == nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "username exists!"})
		return
	}
	curser, err := collection.InsertOne(context.TODO(), User{
		ID:         primitive.NewObjectID(),
		LineId:     newUser.LineId,
		Name:       newUser.Name,
		Points:     newUser.Points,
		PictureUrl: newUser.PictureUrl})
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	id := curser.InsertedID

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "inserted", "id": id})
}

func UpdateUser(c *gin.Context) {
	var updateUser User
	var updateUser2 User
	if err := c.BindJSON(&updateUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no data"})
		return
	}
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.USER)
	err2 := collection.FindOneAndUpdate(context.TODO(),
		bson.M{"lineid": updateUser.LineId},
		bson.M{"$set": updateUser}).Decode(&updateUser2)

	if err2 != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err2.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "updated", "id": err2})

}

func UpdateUserPoint(c *gin.Context) {
	var updateUser User
	var updateUser2 User
	if err := c.BindJSON(&updateUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no data"})
		return
	}
	if updateUser.Points < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "point must be greater than 0"})
		return
	}
	client, err := db.ConnectDatabase()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	collection := client.Database(db.DB).Collection(db.USER)
	err = collection.FindOne(context.TODO(), bson.M{"lineid": updateUser.LineId}).Decode(&updateUser2)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "username does not exist!"})
		return
	}
	point := updateUser2.Points + updateUser.Points
	err = collection.FindOneAndUpdate(context.TODO(),
		bson.M{"lineid": updateUser.LineId}, bson.M{"$set": bson.M{"points": point}}).Decode(&updateUser2)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "update failed!"})
		return
	}
	UpdateStatus(c, client, updateUser.LineId)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "point added", "point": point})
}
