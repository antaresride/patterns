package main

import (
	"fmt"
	"sync"
)

var payment *Payment
var once sync.Once

func GetInstance() *Payment {
	once.Do(func() {
		fmt.Println("Creating unique payment now")
		payment = &Payment{}
	})
	return payment
}
