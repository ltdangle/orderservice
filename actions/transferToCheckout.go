package actions

import (
	"fmt"
	"orders/model/read"
)

type CheckoutUrl string
type CheckoutTransfer struct {
	u CheckoutUrl
}

func NewCheckoutTransfer(url CheckoutUrl) *CheckoutTransfer {
	return &CheckoutTransfer{u: url}
}

func (c *CheckoutTransfer) Url(order *read.Order) string {
	return string(c.u) + fmt.Sprintf("?cart=%s&total=%d", order.Uuid.String, order.Total)
}
