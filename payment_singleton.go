package main

import (
	"fmt"
	"sync"
)

var (
	instance *Payment
	once     sync.Once
)

func GetInstance(merchant string, customer string, amount float64) *Payment {

	once.Do(func() {
		fmt.Println("Creating unique payment now")
		instance = &Payment{
			Merchant: merchant,
			Customer: customer,
			Amount:   amount,
		}
	})
	return instance
}
