package req

type AgentsReq struct {
	Name string `json:"name" validate:"required,min=1,max=10"`
	Role int    `json:"role" validate:"required"`
}
