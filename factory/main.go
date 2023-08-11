package main

import (
	"errors"
	"fmt"
)

type PaymentGatewayType int

const (
	PayPalGateway PaymentGatewayType = iota
	StripeGateway
)

type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

// concrete payment gateway
type PayPal struct{}

func (pg *PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Paypal payment of $%.2f\n", amount)
	return nil
}

type Stripe struct{}

func (sg *Stripe) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	return nil
}

func NewPaymentGateway(gwType PaymentGatewayType) (PaymentGateway, error) {
	switch gwType {
	case PayPalGateway:
		return &PayPal{}, nil
	case StripeGateway:
		return &Stripe{}, nil
	default:
		return nil, errors.New("unsupported payment gateway un")
	}
}

func main() {
	paypalGateway, _ := NewPaymentGateway(PayPalGateway)
	paypalGateway.ProcessPayment(200.00)

	stripeGateway, _ := NewPaymentGateway(StripeGateway)
	stripeGateway.ProcessPayment(599.00)
}
