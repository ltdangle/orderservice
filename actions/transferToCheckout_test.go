package actions_test

import (
	"database/sql"
	"orders/actions"
	"orders/model/read"
	"strconv"
	"testing"
)

func Test_TransferToCheckout(t *testing.T) {
	order := &read.Order{Uuid: sql.NullString{String: "someuuid"}, Total: 100}
	url := "http://checkout.url"
	action := actions.NewCheckoutTransfer(url)
	got := action.Url(order)

	wanted := url + "?cart=" + order.Uuid.String + "&total=" + strconv.Itoa(order.Total)
	if got != wanted {
		t.Fatalf("wanted %s, got %s", wanted, got)
	}
}
