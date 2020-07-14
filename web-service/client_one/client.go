package client_one

import "github.com/DanoFP/dp-golang-training/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
