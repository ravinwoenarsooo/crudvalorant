package entity

type Roles struct {
	Role_ID   uint   `json:"roles_id" gorm:"primarykey"`
	Role_Name string `json:"role_name"`
}
