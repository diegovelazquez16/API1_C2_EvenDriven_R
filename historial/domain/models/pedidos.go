package models


type Historial struct {
    ID            uint    `gorm:"primaryKey"`

    Estado        string  `gorm:"not null"` 

}

    

