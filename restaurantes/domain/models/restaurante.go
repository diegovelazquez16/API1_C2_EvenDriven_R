package models


type Restaurante struct {
    ID         uint    `gorm:"primaryKey"`
    Nombre     string  `gorm:"not null"`
    Direccion  string
    Producto1  string  `gorm:"not null"`
    Precio1    float64 `gorm:"not null"`
    Producto2  string  `gorm:"not null"`
    Precio2    float64 `gorm:"not null"`
    Producto3  string  `gorm:"not null"`
    Precio3    float64 `gorm:"not null"`
}
