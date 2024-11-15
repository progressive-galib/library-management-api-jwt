package models

type Author struct {
    ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Name  string `json:"name" gorm:"not null;unique"`
    Books []Book `json:"books" gorm:"foreignKey:AuthorID"`
}

