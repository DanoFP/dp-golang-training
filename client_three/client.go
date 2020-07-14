package client_three

import (
	"../singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
