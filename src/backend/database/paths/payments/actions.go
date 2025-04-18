package database_payments

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func InitializePaymentTransaction(ctx context.Context, customerID string, productID string, productPrice int64, productDeposit int64, productMode string) (*PaymentTransaction, error) {
	paymentTransactionID := uuid.New().String()

	paymentTransaction := &PaymentTransaction{
		ID:             paymentTransactionID,
		CustomerID:     customerID,
		ProductID:      productID,
		ProductPrice:   productPrice,
		ProductDeposit: productDeposit,
		ProductMode:    productMode,
		Paid:           false,
	}

	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction)
	if err != nil {
		return nil, err
	}
	return paymentTransaction, nil
}

func CompletePaymentTransaction(ctx context.Context, paymentTransactionID string) (*PaymentTransaction, error) {
	paymentTransaction := &PaymentTransaction{}
	paymentTransactionRow := DatabasePayments.Get(ctx, paymentTransactionID)
	if paymentTransactionRow == nil {
		return nil, errors.New("payment transaction not found")
	}

	paymentTransaction.Paid = true
	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction)
	if err != nil {
		return nil, err
	}

	return paymentTransaction, nil
}
