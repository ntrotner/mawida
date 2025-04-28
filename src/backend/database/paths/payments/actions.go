package database_payments

import (
	"context"
	"errors"
	"time"

	"github.com/go-kivik/kivik"
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
		CreatedAt:      time.Now().Unix(),
	}

	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction)
	if err != nil {
		return nil, err
	}
	return paymentTransaction, nil
}

func CompletePaymentTransaction(ctx context.Context, paymentTransactionID string) (*PaymentTransaction, error) {
	paymentTransaction := FindPaymentTransactionById(ctx, paymentTransactionID)
	if paymentTransaction == nil {
		return nil, errors.New("payment transaction not found")
	}

	paymentTransaction.Paid = true
	paymentTransaction.PaidAt = time.Now().Unix()
	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction, kivik.Options{
		"_rev": paymentTransaction.Rev,
	})
	if err != nil {
		return nil, err
	}

	return paymentTransaction, nil
}

func ReturnPaymentTransaction(ctx context.Context, paymentTransactionID string) (*PaymentTransaction, error) {
	paymentTransaction := FindPaymentTransactionById(ctx, paymentTransactionID)
	if paymentTransaction == nil {
		return nil, errors.New("payment transaction not found")
	}

	paymentTransaction.ReturnedAt = time.Now().Unix()
	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction, kivik.Options{
		"_rev": paymentTransaction.Rev,
	})
	if err != nil {
		return nil, err
	}

	return paymentTransaction, nil
}

func RefundDepositPaymentTransaction(ctx context.Context, paymentTransactionID string, depositRefundedAmount int64) (*PaymentTransaction, error) {
	paymentTransaction := FindPaymentTransactionById(ctx, paymentTransactionID)
	if paymentTransaction == nil {
		return nil, errors.New("payment transaction not found")
	}

	paymentTransaction.DepositRefundedAmount = depositRefundedAmount
	paymentTransaction.DepositRefundedAt = time.Now().Unix()
	_, err := DatabasePayments.Put(ctx, paymentTransaction.ID, paymentTransaction, kivik.Options{
		"_rev": paymentTransaction.Rev,
	})
	if err != nil {
		return nil, err
	}

	return paymentTransaction, nil
}
