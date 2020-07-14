package persistence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DanoFP/dp-golang-training/dt"
	"github.com/gorilla/mux"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(dt.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range dt.Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article dt.Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	dt.Articles = append(dt.Articles, article)

	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range dt.Articles {
		if article.Id == id {
			dt.Articles = append(dt.Articles[:index], dt.Articles[index+1:]...)
		}
	}

}
