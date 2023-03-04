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

	client, err :=mongo.Connect(context.TODO(),clientOptions)

	if err !=nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)

	if err !=nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance created")

}

//Get all the task route
func GetAllTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	payload := getAlltask()
	json.NewEncoder(w).Encode(payload)
}

//CreateTask create task route

func CreateTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	var task models.ToDoList 
	_=json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

//TaskComplete Update task route

func TaskComplete(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	
	params:=mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}







