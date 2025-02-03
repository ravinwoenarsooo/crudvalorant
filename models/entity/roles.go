package entity

type Roles struct {
	Role_Id   uint   `json:"id" gorm:"primaryKey"`
	Role_Name string `json:"name"`
}
