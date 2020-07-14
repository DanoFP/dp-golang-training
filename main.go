package main

import (
	"fmt"
	"sync"

	"./client_one"
	"./client_two"
	"./singleton"
)

//Local Server
/* func main() {
	dt.Articles = []dt.Article{
		dt.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		dt.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	server.HandleRequests()
} */

//Singleton test
func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
