func (s *CardService) GetCards(query *models.CardQuery) ([]*models.Card, int64, error) {
	db := utils.DB.Model(&models.Card{}).Where("merchant_id = ?", query.MerchantID)
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	var cards []*models.Card
	if err := db.Offset((query.Page-1)*query.PageSize).Limit(query.PageSize).Find(&cards).Error; err != nil {
		return nil, 0, err
	}
	return cards, total, nil
}