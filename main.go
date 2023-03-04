package main

import(
	"fmt"
	"log"
	"net/http"
	"Todo-app-golang/router"
)

func main(){
	r:= router.Router()
	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000",r))
}