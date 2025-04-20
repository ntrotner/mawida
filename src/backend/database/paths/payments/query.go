package database_payments

import (
	"context"
	database_common "template_backend/database/common"
)

func createFindPaymentTransactionQuery(id *string, fields []interface{}) database_common.Query {
	query := database_common.Query{
		Selector: map[string]interface{}{},
		Fields:   fields,
		Limit:    1,
	}

	if id != nil {
		query.Selector["_id"] = map[string]interface{}{"$eq": *id}
	}

	return query
}

func FindPaymentTransactionById(ctx context.Context, id string) *PaymentTransaction {
	if DatabasePayments == nil {
		return nil
	}

	query := createFindPaymentTransactionQuery(&id, []interface{}{
		"_id",
		"_rev",
		"customerId",
		"productId",
		"productPrice",
		"productDeposit",
		"productMode",
		"paid",
		"createdAt",
		"paidAt",
		"returnedAt",
		"depositRefundedAmount",
		"depositRefundedAt",
	})
	rows, err := DatabasePayments.Find(ctx, query)
	if err != nil {
		return nil
	}

	var paymentTransaction PaymentTransaction
	if rows.Next() {
		if err := rows.ScanDoc(&paymentTransaction); err != nil {
			return nil
		}
		return &paymentTransaction
	}

	return nil
}

func GetAllPaymentTransactions(ctx context.Context) []*PaymentTransaction {
	if DatabasePayments == nil {
		return nil
	}

	query := createFindPaymentTransactionQuery(nil, []interface{}{"_id", "customerId", "productId", "productPrice", "productDeposit", "productMode", "paid", "createdAt", "paidAt", "returnedAt", "depositRefundedAmount", "depositRefundedAt"})
	query.Limit = 5000
	rows, err := DatabasePayments.Find(ctx, query)
	if err != nil {
		return nil
	}

	var paymentTransactions []*PaymentTransaction
	for rows.Next() {
		var paymentTransaction PaymentTransaction
		if err := rows.ScanDoc(&paymentTransaction); err != nil {
			return nil
		}
		paymentTransactions = append(paymentTransactions, &paymentTransaction)
	}

	return paymentTransactions
}
