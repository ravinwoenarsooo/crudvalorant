package entity

type Roles struct {
	Role_Id   uint     `json:"id" gorm:"primaryKey"`
	Role_Name string   `json:"name"`
	Agents    []Agents `json:"agents" gorm:"foreignKey:Role_Id"` // One-to-Many Relationship
}
