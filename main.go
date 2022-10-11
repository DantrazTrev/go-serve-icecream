package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

const _PORT = ":8080"

type Item struct {
	Name string
	Price string
	Address string
}
type PageVariables struct {
	PageTitle string
	Items []Item
}

var items []Item

func main(){
	http.HandleFunc("/",landing)
	http.HandleFunc("/items/",getItems)
	http.HandleFunc("/add-item/",addItem)


	fmt.Println("Server is running on",_PORT)
	log.Fatal(http.ListenAndServe(_PORT,nil))
}

func addItem(w http.ResponseWriter,req *http.Request) {
	err:=req.ParseForm()
	if err !=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		log.Print("Form Parsing Error")

	}
	fmt.Println()
	item := Item{
		Name: req.FormValue("item"),
		Price: req.FormValue("price"),
		Address: req.FormValue("address"),

	}
	items=append(items,item)
	http.Redirect(w,req,"/items/",http.StatusSeeOther)
}
func landing(w http.ResponseWriter,req *http.Request) {
	fmt.Println("Welcome to the ICE CREAM STORE")

}

func getItems(w http.ResponseWriter,req *http.Request) {
	pageVariables:=PageVariables{
		PageTitle : "YSISIC",
		Items: items,
	}

	t,err:=template.ParseFiles("index.html")
	if err!=nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		log.Print("Template Parsing Error")
	}
	err=t.Execute(w,pageVariables);
	fmt.Println(t,err)
}