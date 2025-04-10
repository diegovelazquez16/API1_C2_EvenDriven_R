package repository

import (
	"gorm.io/gorm"

	"holamundo/restaurantes/domain/models"
)


type RestauranteRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}
func (r *RestauranteRepositoryImpl) Create(restaurante *models.Restaurante) error {
	return r.DB.Create(restaurante).Error
}

func (r *RestauranteRepositoryImpl) GetAll() ([]models.Restaurante, error) {
	var restaurantes []models.Restaurante
	err := r.DB.Find(&restaurantes).Error
	return restaurantes, err
}

func (r *RestauranteRepositoryImpl) GetByID(id uint) (*models.Restaurante, error) {
	var restaurante models.Restaurante
	err := r.DB.First(&restaurante, id).Error
	return &restaurante, err
}

func (r *RestauranteRepositoryImpl) Update(restaurante *models.Restaurante) error{
	return r.DB.Save(restaurante).Error
}

func (r *RestauranteRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.Restaurante{}, id).Error
} 