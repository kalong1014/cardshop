func (s *PaymentService) ApplyRefund(orderID uint, userID uint, reason string) error {
	var order models.Order
	if err := utils.DB.First(&order, orderID).Error; err != nil {
		return err
	}
	if order.Status != models.OrderStatusPaid {
		return errors.New("订单未支付，无法退款")
	}
	
	return utils.DB.Create(&models.Refund{
		OrderID:    orderID,
		UserID:     userID,
		MerchantID: order.MerchantID,
		Amount:     order.Amount,
		Reason:     reason,
	}).Error
}

func (s *PaymentService) ApproveRefund(refundID uint) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		var refund models.Refund
		if err := tx.First(&refund, refundID).Error; err != nil {
			return err
		}
		// 更新退款状态
		if err := tx.Model(&refund).Update("status", models.RefundStatusApproved).Error; err != nil {
			return err
		}
		// 恢复卡密状态（假设卡密核销后可退款）
		var order models.Order
		if err := tx.First(&order, refund.OrderID).Error; err != nil {
			return err
		}
		return tx.Model(&models.Card{}).Where("id = ?", order.CardID).Update("status", models.CardStatusUnused).Error
	})
}

// service/payment_service.go
func (s *PaymentService) DailyReconciliation() {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
			var orders []models.Order
			utils.DB.
				Where("status = ? AND created_at >= ?", models.OrderStatusPaid, yesterday).
				Find(&orders)
			
			// 模拟获取支付渠道对账单（需对接实际渠道API）
			bill := s.mockChannelBill(yesterday)
			orderTotal := calculateOrderTotal(orders)
			billTotal := bill.Amount

			if orderTotal != billTotal {
				utils.DB.Create(&models.ReconciliationLog{
					Date:      yesterday,
					OrderTotal: orderTotal,
					BillTotal: billTotal,
					Difference: billTotal - orderTotal,
					Status:    models.ReconciliationStatusMismatch,
				})
			}
		}
	}()
}

// 模拟对账单
func (s *PaymentService) mockChannelBill(date string) models.PaymentBill {
	return models.PaymentBill{
		Date:   date,
		Amount: 1000.00, // 假设渠道对账单金额
	}
}