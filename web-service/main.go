package main

import (
	"github.com/DanoFP/dp-golang-training/web-service/dt"
	"github.com/DanoFP/dp-golang-training/web-service/server"
)

//Local Server
func main() {
	dt.Articles = []dt.Article{
		dt.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		dt.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	server.HandleRequests()
}
