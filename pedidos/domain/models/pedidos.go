package models



type Pedido struct {
    ID            uint    `gorm:"primaryKey"`
    UsuarioID     uint    `gorm:"not null"`
    RestauranteID uint    `gorm:"not null"`
    Producto1     string
    Producto2     string
    Producto3     string
    Total         float64 `gorm:"not null"`
    Estado        string  `gorm:"not null"` 
    PagoID        uint    `gorm:"not null;default:1"`
}

    

