package service

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (s *Service) deleteRecords(tables []schema.Tabler, tx *gorm.DB, relyColumName string, relyValue string) error {
	if relyColumName == "" || relyValue == "" {
		return errors.New("missing delete rely")
	}
	for _, table := range tables {
		if err := tx.Where(relyColumName+" = ?", relyValue).Delete(table).Error; err != nil {
			return err
		}
	}
	return nil
}
