package actions_test

import (
	// "orders/actions"
	"errors"
	"orders/actions"
	"orders/model/write"
	"testing"
)

func Test_CreateOrder_Success(t *testing.T) {
	// SaveOrder function type mock.
	saveOrder:=func(order *write.Order) error {
		return nil
	}

	action:=actions.NewCreateOrder(saveOrder)
	_,err:=action.Create(actions.NewOrderRequest{})

	if err!=nil{
		t.Logf("Cannot create new order")
	}
}

func Test_CreateOrder_Failure(t *testing.T) {
	// SaveOrder function type mock.
	saveOrder:=func(order *write.Order) error {
		return errors.New("could not save order")
	}

	action:=actions.NewCreateOrder(saveOrder)
	_,err:=action.Create(actions.NewOrderRequest{})

	if err==nil{
		t.Logf("save order error not handled")
	}
}
