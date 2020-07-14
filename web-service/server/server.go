package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DanoFP/dp-golang-training/web-service/persistence"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", persistence.ReturnAllArticles)
	myRouter.HandleFunc("/article", persistence.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", persistence.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", persistence.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

/*
ADD
fetch("/article",{method:"POST",body:JSON.stringify({
    "Id": "3",
    "Title": "Created Post",
    "desc": "The description for my new post",
    "content": "my articles content"
})}).then((res)=>console.log(res.json()));

DELETE
fetch("/article/3",{method:"DELETE"}).then((res)=>console.log(res.json()));
*/
