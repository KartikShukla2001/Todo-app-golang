package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

//DB connection String
const connectionString = "mongodb://localhost:27017"

const dbName = "test"

const collName ="todoList"

var collection *mongo.Collection

func init(){

	clientOptions :=options.Client().ApplyURI(connectionString)
}

