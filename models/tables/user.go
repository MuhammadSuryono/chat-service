package tables

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhotoPath string `json:"photo_path"`
}
