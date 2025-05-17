func (s *OrderService) GetOrders(query *models.OrderQuery) ([]*models.Order, int64, error) {
	db := utils.DB.Model(&models.Order{})
	if query.MerchantID > 0 {
		db = db.Where("merchant_id = ?", query.MerchantID)
	}
	if query.UserID > 0 {
		db = db.Where("user_id = ?", query.UserID)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	var orders []*models.Order
	if err := db.
		Preload("Card"). // 预加载卡密信息
		Offset((query.Page-1)*query.PageSize).
		Limit(query.PageSize).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}