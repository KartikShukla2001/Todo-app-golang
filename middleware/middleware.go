package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"Todo-app-golang/models"


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

	collection = client.Database(dbName).Collection(collName)

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

func UndoTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")

	params:=mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params:=mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}


func DeleteAllTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count:=deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

func getAlltask() []primitive.M{
	cur,err :=collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}

	var results[]primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e:=cur.Decode(&result)
		if e !=nil{
			log.Fatal(err)
		}
		results=append(results,result)
	}

	if err := cur.Err(); err !=nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results;

}

func undoTask(task string){
	fmt.Println(task)
	id,_ :=primitive.ObjectIDFromHex(task)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set": bson.M{"status":false}}
	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("modified count", result.ModifiedCount)
}

func insertOneTask(task models.ToDoList){

	insertResult,err:=collection.InsertOne(context.Background(),task)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record",insertResult.InsertedID)
}

func taskComplete(task string){
	fmt.Println(task)
		id,_ := primitive.ObjectIDFromHex(task)
		filter:=bson.M{"_id": id}
		update:=bson.M{"$set": bson.M{"status":true}}
		result,err:=collection.UpdateOne(context.Background(),filter,update)
		if err !=nil{
			log.Fatal(err)
		}

		fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneTask(task string) {
	fmt.Println(task)
	id,_:=primitive.ObjectIDFromHex(task)
	filter :=bson.M{"_id:": id}
	d,err :=collection.DeleteOne(context.Background(),filter)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

func deleteAllTask() int64{
	d,err :=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted Document",d.DeletedCount)
	return d.DeletedCount
}








