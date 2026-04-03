package main

import (
	"fmt"
)

type Authorization struct {
	cryptography string
}

// The complex object
type Payment struct {
	Merchant      string
	Customer      string
	Amount        float64
	Authorization Authorization
}
type PaymentBuilder struct {
	pay Payment
}

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{
		pay: Payment{
			Merchant: "default-merchant",
			Customer: "default-customer",
			Amount:   0.00,
		},
	}
}

func (pb *PaymentBuilder) WithMerchant(m string) *PaymentBuilder {
	pb.pay.Merchant = m
	return pb
}

func (pb *PaymentBuilder) WithCustomer(c string) *PaymentBuilder {
	pb.pay.Customer = c
	return pb
}

func (pb *PaymentBuilder) WithAmount(a float64) *PaymentBuilder {
	pb.pay.Amount = a
	return pb
}

func (pb *PaymentBuilder) WithAuthorization(auth Authorization) *PaymentBuilder {
	pb.pay.Authorization = auth
	return pb
}

func (pb *PaymentBuilder) Build() Payment {
	return pb.pay
}

func (l *Payment) Pay() {
	fmt.Println("I am the legacy method")
}
func main() {
	// Builder
	payment := NewPaymentBuilder().
		WithMerchant("Tesco").
		WithCustomer("Alice").
		WithAmount(13.00).
		WithAuthorization(Authorization{cryptography: "ecdsa"}).
		Build()
	fmt.Printf("%+v\n", payment)

	// Adapter
	legacyPayment := &payment
	adapterPayment := &PaymentAdapter{Payment: legacyPayment}
	Pay(adapterPayment)
	payment.Pay()

	// Singleton
	GetInstance("Louie", "Maria", 10.00)

	GetInstance("Fernand", "Luca", 30.00)

	GetInstance("Sheila", "Jorge", 20.00)

}
