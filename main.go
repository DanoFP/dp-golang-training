package main

import (
	"fmt"
	"os"

	"github.com/DanoFP/dp-golang-training/factory"
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
/* func main() {
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
} */
//DB Connection
func main() {
	var t int
	fmt.Print("Select connection : (1:mysql 2:postgres) ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error to read option: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("invalid db")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Printf("something went wrong %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))

	err = conn.Close()
	if err != nil {
		fmt.Printf("something went wrong cant close connection %v", err)
	}
}
