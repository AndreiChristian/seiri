package models


type Todo struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string `gorm:"not null"`
    Description string
   Completed   bool   `gorm:"default:false"`
    UserID      uint   `gorm:"not null"`
    User        User   `gorm:"foreignKey:UserID"`
}

