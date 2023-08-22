package actions

import (
	"fmt"
	"orders/model/read"
)

type CheckoutTransfer struct {
	u string
}

func NewCheckoutTransfer(url string) *CheckoutTransfer {
	return &CheckoutTransfer{u: url}
}

func (c *CheckoutTransfer) Url(order *read.Order) string {
	return c.u + fmt.Sprintf("?cart=%s&total=%d", order.Uuid.String, order.Total)
}
