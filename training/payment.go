package training

import "fmt"

type Payment interface {
	Pay(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

type PayPal struct {
	Email string
}

type PayLib struct {
	PhoneNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card ending in %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])

}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, pp.Email)
}

func (pn PayLib) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, pn.PhoneNumber)
}

func ProcessPayment(p Payment, amount float64) {
	fmt.Println(p.Pay(amount))
}
