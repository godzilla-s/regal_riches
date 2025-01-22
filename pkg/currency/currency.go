package currency

import "fmt"

type TON struct {
	amount int64
}

type RR struct {
	amount int64
}

func ToTON(amount int64) TON {
	return TON{amount: amount}
}

func ToRR(amount int64) RR {
	return RR{amount: amount}
}

func (t TON) ToRR() RR {
	return RR{
		amount: t.amount * 10000,
	}
}

func (t TON) String() string {
	return fmt.Sprintf("%.2f TON", float64(t.amount/10000.0))
}

func (r RR) ToTON() TON {
	return TON{amount: r.amount / 10000}
}
