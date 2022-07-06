package main


import (
	"fmt"
	"kano-library/routes"
	"net/http"
	// "kano-library/crud"
)

type Books struct {
	Title 		string `json:"title"`
	Author 		string `json:"author"`
	Synopsis 	string `json:"synopsis"`
	Release 	string `json:"release"`
}

func main() {
	// crud.Insert("books", Books{Title: "Alaia", Author: "Raden", Synopsis: "opqrs", Release: "2020"})
	// crud.Find("books", Books{})
	// crud.Update()
	// crud.Remove()


	http.HandleFunc("/books", routes.FindAll)
	http.HandleFunc("/books/find", routes.FindByIdBook)
	http.HandleFunc("/books/insert", routes.InsertBook)
	http.HandleFunc("/books/update", routes.UpdateBook)
	http.HandleFunc("/books/delete", routes.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}