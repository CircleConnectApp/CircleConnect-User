package models

type User struct {
    ID             uint   `json:"id" gorm:"primaryKey"`
    Email          string `json:"email"`
    GoogleSub      string `json:"google_sub"`
    Name           string `json:"name"`
    Role           string `json:"role"`
    Birthday       string `json:"birthday"`
    Bio            string `json:"bio"`
    InstagramURL   string `json:"instagram_url"`
    ProfilePicture string `json:"profile_picture"`
    Location       string `json:"location"`
}
