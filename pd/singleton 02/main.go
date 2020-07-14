package singleton

import (
	"fmt"
	"sync"

	"github.com/DanoFP/dp-golang-training/singleton"
	"github.com/DanoFP/dp-golang-training/web-service/client_one"
	"github.com/DanoFP/dp-golang-training/web-service/client_two"
)

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
