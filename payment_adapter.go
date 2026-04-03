package main

import "fmt"

type ModernPayment interface {
	Pay() float64
}

type PaymentAdapter struct {
	Payment *Payment
}

func (a *PaymentAdapter) Pay() float64 {
	return a.Payment.Amount
}

func Pay(p ModernPayment) {
	fmt.Println("I am the modern payment method")
	fmt.Printf("%+v\n", p.Pay())

}
