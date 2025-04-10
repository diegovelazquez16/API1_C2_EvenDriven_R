package repository

import (
	"gorm.io/gorm"

	"holamundo/historial/domain/models"
)


type HistorialRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}
func (r *HistorialRepositoryImpl) Create(historial *models.Historial) error {
	return r.DB.Create(historial).Error
}

func (r *HistorialRepositoryImpl) GetAll() ([]models.Historial, error) {
	var historiales []models.Historial
	err := r.DB.Find(&historiales).Error
	return historiales, err
}

func (r *HistorialRepositoryImpl) GetByID(id uint) (*models.Historial, error) {
	var historial models.Historial
	err := r.DB.First(&historial, id).Error
	return &historial, err
}

func (r *HistorialRepositoryImpl) Update(historial *models.Historial) error{
	return r.DB.Save(historial).Error
}

func (r *HistorialRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.Historial{}, id).Error
} 