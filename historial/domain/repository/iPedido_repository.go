package repository

import (
	"holamundo/historial/domain/models"
)

type IHistorialRepository interface {
	Create(historial *models.Historial) error
	GetAll() ([]models.Historial, error)
	GetByID(id uint) (*models.Historial, error)
	Update(historial *models.Historial) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}
