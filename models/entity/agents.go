package entity

type Agents struct {
	Agent_Id   uint   `json:"id" gorm:"primaryKey"`
	Agent_Name string `json:"name"`
	Role_Id    int    `json:"role_id"` // Foreign Key
	Role       Roles  `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
