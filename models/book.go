package models

type Book struct {
    ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Title     string `json:"title" gorm:"not null"`
    AuthorID  int    `json:"author_id" gorm:"not null"`
    Available bool   `json:"available" gorm:"default:true"`
    Author    Author `json:"author" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

