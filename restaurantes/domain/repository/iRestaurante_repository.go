package repository

import (
	"holamundo/restaurantes/domain/models"
)

type IRestauranteRepository interface {
	Create(restaurante *models.Restaurante) error
	GetAll() ([]models.Restaurante, error)
	GetByID(id uint) (*models.Restaurante, error)
	Update(restaurante *models.Restaurante) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}
