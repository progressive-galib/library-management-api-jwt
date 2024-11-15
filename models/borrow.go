package models

import "time"

type Borrow struct {
    ID       int       `json:"id" gorm:"primaryKey;autoIncrement"`
    BookID   int       `json:"book_id" gorm:"not null"`
    UserID   int       `json:"user_id" gorm:"not null"`
    DueDate  time.Time `json:"due_date" gorm:"not null"`
    Book     Book      `json:"book" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    User     User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

