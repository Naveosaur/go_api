package dto

import "mime/multipart"

type PostResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"-"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	Post       string `json:"post"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
}

type PostRequest struct {
	UserID  int    `form:"user_id"` // Karena ada pic_url jadi tidak bisa melalui json
	Post    string `form:"post"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}