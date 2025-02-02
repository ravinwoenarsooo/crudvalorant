package entity

// Ini buat tabel di database buat di migrate.
type Agent struct {
	Agent_ID   uint   `json:"id" gorm:"primarykey"`
	Agent_Name string `json:"name"`
	Role_ID    uint   `json:"-"`
	Role       Roles  `json:"role" gorm:"foreignkey:Role_ID"`
}
