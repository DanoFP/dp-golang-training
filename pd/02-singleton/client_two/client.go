package client_two

import "github.com/DanoFP/dp-golang-training/pd/02-singleton/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
