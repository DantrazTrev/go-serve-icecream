package main

import (
	"fmt"
	"net/http"
	"log"
)

const _PORT = ":8080"

func main(){
	http.HandleFunc("/",landing)
	fmt.Println("Server is running on",_PORT)
	log.Fatal(http.ListenAndServe(_PORT,nil))
}

func landing(w http.ResponseWriter,req *http.Request) {
	fmt.Println("Welcome to the ICE CREAM STORE")

}