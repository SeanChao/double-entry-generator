package trading212

import (
	"log"
	"strconv"
	"time"
)

func (a *Trading212) translateToOrdersFromMap(m map[string]string) error {
	var order Order
	var err error

	order.Action = m["Action"]
	order.Time, err = time.Parse(localTimeFmt, m["Time"]+" +0000 UTC")
	if err != nil {
		log.Println("parse time error:", m["Time"], err)
		return err
	}
	// The money is the change of the account, spending is negative, deposit is positive.
	money, err := strconv.ParseFloat(m["Total"], 32)
	// order.Money = math.Abs(money)
	order.Money = money
	if err != nil {
		log.Println("parse money error:", m["Total"], err)
		return err
	}
	// Use "Currency", if not available, use "Currency (Original)"
	order.Currency = m["Currency"]
	if order.Currency == "" {
		order.Currency = m["Currency (Total)"]
	}
	order.Notes = m["Notes"]
	order.ID = m["ID"]
	order.Peer = m["Merchant name"]
	order.MerchantCategory = m["Merchant category"]

	a.Orders = append(a.Orders, order)
	return nil
}
