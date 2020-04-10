package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct
type User struct {
	ID        primitive.ObjectID `json:"id" binding:"omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" binding:"omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastName" binding:"omitempty" bson:"lastname,omitempty"`
	Age       int                `json:"age,omitempty" binding:"omitempty" bson:"age,omitempty"`
}

// Response struct
type Res struct {
	Ctx     *gin.Context	`json:"-"`
	Err		error			`json:"error,omitempty"`
	Msg		string			`json:"message,omitempty"`
	Status  int				`json:"status,omitempty"`
	Data	interface{}		`json:"data,omitempty"`
}

var db *mongo.Collection

func dbConnect() *mongo.Collection {
	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Fatal("db env not set")
	}
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the db...")
	return client.Database("go-testing").Collection("users")
}

func Router() *gin.Engine {
	db = dbConnect()
	router := gin.New()
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", checkUser, getUser)
	router.POST("/users", createUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	return router
}

// middleware for checking if user exists
func checkUser(c *gin.Context) {
	var user User
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		resErr(Res{Ctx: c, Err: err})
		return
	}
	c.Set("user", user)
}

// Get single user
func getUser(c *gin.Context) {
	user := c.MustGet("user").(User)
	// var user User
	// oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	// err := db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
	// if err != nil {
	// 	resErr(Res{Ctx: c, Err: err})
	// 	return
	// }
	res(Res{Ctx: c, Data: user})
}

// Get all users
func getAllUsers(c *gin.Context) {
	var users []User
	cur, err := db.Find(context.TODO(), bson.D{{}},)
	if err != nil {
    	resErr(Res{Ctx: c, Err: err})
		return
	}
	for cur.Next(context.TODO()) {
		var user User
		_ = cur.Decode(&user)
		users = append(users, user)
	}
	res(Res{Ctx: c, Data: users})
}

// Create single user
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr(Res{Ctx: c, Msg: "Invalid request"})
		return
	}
	if _, err := db.InsertOne(context.TODO(), user); err != nil {
		resErr(Res{Ctx: c })
		return
	}
	res(Res{Ctx: c})
}

// Update single user
func updateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr(Res{Ctx: c, Err: err})
		return
	}
	data, _ := bson.Marshal(user)
	var updateQuery bson.M
	_ = bson.Unmarshal(data, &updateQuery)
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	result, err := db.UpdateOne(
		context.TODO(),
		bson.M{"_id": oid},
		bson.D{{"$set", updateQuery}},
	)
	if err != nil {
		resErr(Res{Ctx: c, Err: err})
		return
	}
	res(Res{Ctx: c, Data: result})
}

// delete single user
func deleteUser(c *gin.Context) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if _, err := db.DeleteOne(context.TODO(), bson.M{"_id": oid}); err != nil {
		resErr(Res{Ctx: c, Err: err})
		return
	}
	res(Res{Ctx: c })
}

// sends a json response if there is an error
func resErr(r Res) {
	if r.Msg == "" {
		r.Msg = "Error occured"
	}
	if r.Status == 0 {
		r.Status = 400
	}
	log.Println("error: ", r.Err)
	r.Ctx.JSON(r.Status, r)
}

// sends a json resonse on success
func res(r Res) {
	if r.Msg == "" {
		r.Msg = "Successful!"
	}
	if r.Status == 0 {
		r.Status = 200
	}
	r.Ctx.JSON(r.Status, r)
}
