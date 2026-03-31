package main

import "fmt"

type Authorization struct {
	cryptography string
}

// the complex object
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

func (pb *PaymentBuilder) WithAmmount(a float64) *PaymentBuilder {
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

func main() {
	payment := NewPaymentBuilder().
		WithMerchant("Tesco").
		WithCustomer("Alice").
		WithAmmount(13.00).
		WithAuthorization(Authorization{cryptography: "edca"}).
		Build()
	fmt.Println(payment)

}
