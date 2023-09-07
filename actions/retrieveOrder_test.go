package actions_test

import (
	"errors"
	"orders/actions"
	"orders/model/read"
	"reflect"
	"testing"
)

func Test_RetrieveOrder(t *testing.T) {
	order := &read.Order{}

	tests := []struct {
		testName    string
		findOrder   read.FindOrder
		orderResult *read.Order
		errResult   error
	}{
		{"success",
			func(uuid string) (*read.Order, error) {
				return order, nil
			},
			order,
			nil,
		},
		{"failure - order finder error",
			func(uuid string) (*read.Order, error) {
				return nil, errors.New("order not found")
			},
			nil,
			errors.New("order not found"),
		},
		{"failure - order is nil",
			func(uuid string) (*read.Order, error) {
				return nil, nil
			},
			nil,
			errors.New("order is nil"),
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			action := actions.NewRetrieveOrder(tt.findOrder)
			order, err := action.Retrieve("someuuid")
			if order != tt.orderResult || reflect.TypeOf(err) != reflect.TypeOf(tt.errResult) {
				t.Errorf(tt.testName, order, err, tt.orderResult, tt.errResult)
			}
		})
	}
}
