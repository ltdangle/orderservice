package actions

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"orders/model/write"
	"time"
)

type ConfirmPayment struct {
	orderSaver write.IOrderSaver
	db         *gorm.DB
}

func NewConfirmPayment(orderSaver write.IOrderSaver, db *gorm.DB) *ConfirmPayment {
	return &ConfirmPayment{orderSaver: orderSaver, db: db}
}

func (a *ConfirmPayment) Confirm(orderUuid string, paymentUuid string) error {
	order, err := a.orderSaver.FindById(orderUuid)
	if err != nil {
		return err
	}

	// Retrieve payment.
	var payment write.Payment
	payment.Uuid = uuid.New().String()
	payment.PaymentId = paymentUuid
	payment.OrderId = orderUuid
	payment.Date = time.Now()
	result := a.db.Create(payment)
	if result.Error != nil {
		return err
	}

	// Update payment.
	payment.PaymentId = paymentUuid
	payment.OrderId = orderUuid
	result = a.db.Save(payment)
	if result.Error != nil {
		return result.Error
	}

	// Update order.
	order.Status = "paid"
	order.PaymentId = paymentUuid

	err = a.orderSaver.Update(order)
	if err != nil {
		return err
	}

	return nil
}
