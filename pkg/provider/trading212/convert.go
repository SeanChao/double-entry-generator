package trading212

import (
	"fmt"

	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// convertToIR convert Trading212 bills to IR.
func (a *Trading212) convertToIR() *ir.IR {
	i := ir.New()
	for _, o := range a.Orders {

		irO := ir.Order{
			OrderType: ir.OrderTypeFx,
			Peer:         o.Peer,
			Item:         "",
			Category:     o.MerchantCategory,
			Method:       "",
			PayTime:      o.Time,
			Money:        o.Money,
			OrderID:      &o.ID,
			Type:         convertType(o.Action),
			TypeOriginal: o.Action,
			Note:         o.Notes,
			Units: map[ir.Unit]string{
				ir.TargetUnit: o.Currency,
				ir.BaseUnit:   o.Currency,
			},
			Metadata: map[string]string{
				"id": o.ID,
			},
		}
		if irO.Type == ir.TypeSend {
			irO.Money = -irO.Money
		}
		i.Orders = append(i.Orders, irO)
	}
	return i
}

func convertType(t string) ir.Type {
	switch t {
	case "Card debit":
		return ir.TypeSend
	case "Deposit":
		return ir.TypeRecv
	case "Spending cashback":
		return ir.TypeRecv
	case "Interest on cash":
		return ir.TypeRecv
	default:
		return ir.TypeUnknown
	}
}

// getMetadata get the metadata (e.g. status, method, category and so on.)
//
//	from order.
/* func getMetadata(o Order) map[string]string {
	// FIXME(TripleZ): hard-coded, bad pattern
	source := "支付宝"
	var status, method, category, typeOriginal, orderId, merchantId, paytime string

	paytime = o.PayTime.Format(localTimeFmt)

	if o.DealNo != "" {
		orderId = o.DealNo
	}

	if o.MerchantId != "" {
		merchantId = o.MerchantId
	}

	if o.Category != "" {
		category = o.Category
	}

	if o.TypeOriginal != "" {
		typeOriginal = o.TypeOriginal
	}

	if o.Method != "" {
		method = o.Method
	}

	if o.Status != "" {
		status = o.Status
	}

	return map[string]string{
		"source":     source,
		"payTime":    paytime,
		"orderId":    orderId,
		"merchantId": merchantId,
		"type":       typeOriginal,
		"category":   category,
		"method":     method,
		"status":     status,
	}
} */
