package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"orders/model/read"
	"orders/rest"
	"testing"
)

func Test_RetrieveOrder(t *testing.T) {

	tests := []struct {
		name           string
		findOrder      read.FindOrder
		wantHttpStatus int
	}{
		{"success",
			func(uuid string) (*read.Order, error) {
				return &read.Order{}, nil
			},
			http.StatusOK,
		},
		{"failure - order finder error",
			func(uuid string) (*read.Order, error) {
				return nil, errors.New("order not found")
			},
			http.StatusInternalServerError,
		},
		{"failure - order is nil",
			func(uuid string) (*read.Order, error) {
				return nil, nil
			},
			http.StatusInternalServerError,
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init http testing server.
			r := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
			w := httptest.NewRecorder()
			// Init our controller.
			rspndr := rest.NewResponder("2006-01-02 15:04:05")
			cntrlr := rest.NewRetrieveOrder(tt.findOrder, rspndr)
			cntrlr.Retrieve(w, r)

			gotHttpStatus := w.Result().StatusCode
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf(tt.name, gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
