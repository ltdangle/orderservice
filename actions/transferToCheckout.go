package actions

type CheckoutTransfer struct {
	u string
}

func NewCheckoutTransfer(url string) *CheckoutTransfer {
	return &CheckoutTransfer{u: url}
}

func (c *CheckoutTransfer) Url(orderUuid string) string {
	return c.u + "?cart="+orderUuid
}

