package repository

import (
	"gorm.io/gorm"

	"holamundo/pedidos/domain/models"
)


type PedidosRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}
func (r *PedidosRepositoryImpl) Create(pedidos *models.Pedido) error {
	return r.DB.Create(pedidos).Error
}

func (r *PedidosRepositoryImpl) GetAll() ([]models.Pedido, error) {
	var pedidos []models.Pedido
	err := r.DB.Find(&pedidos).Error
	return pedidos, err
}

func (r *PedidosRepositoryImpl) GetByID(id uint) (*models.Pedido, error) {
	var pedido models.Pedido
	err := r.DB.First(&pedido, id).Error
	return &pedido, err
}

func (r *PedidosRepositoryImpl) Update(pedido *models.Pedido) error{
	return r.DB.Save(pedido).Error
}

func (r *PedidosRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.Pedido{}, id).Error
} 