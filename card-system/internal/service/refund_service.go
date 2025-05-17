package service

import (
	"card-system/internal/model"
)

func (s *RefundService) ProcessRefund(refundID uint) error {
	tx := utils.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var refund model.Refund
	if err := tx.First(&refund, refundID).Error; err != nil {
		return err
	}

	var order model.Order
	if err := tx.First(&order, refund.OrderID).Error; err != nil {
		return err
	}

	// 执行退款逻辑（对接支付渠道API）
	if err := s.channel.Refund(order.TransactionID, refund.Amount); err != nil {
		return tx.Rollback().Error
	}

	// 更新订单和卡密状态
	if err := tx.Model(&order).Update("status", model.OrderStatusRefunded).Error; err != nil {
		return tx.Rollback().Error
	}
	if err := tx.Model(&model.Card{}).Where("id = ?", order.CardID).Update("status", model.CardStatusRefunded).Error; err != nil {
		return tx.Rollback().Error
	}

	refund.Status = model.RefundStatusApproved
	if err := tx.Save(&refund).Error; err != nil {
		return tx.Rollback().Error
	}

	tx.Commit()
	order.SendRefundNotification() // 触发通知
	return nil
}
