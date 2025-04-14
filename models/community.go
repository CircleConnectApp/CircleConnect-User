package models

type Community struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Language    string `json:"language"`
    Description string `json:"description"`
}
