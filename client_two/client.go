package client_two

import "github.com/DanoFP/dp-golang-training/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
