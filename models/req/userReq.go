package req

type AgentReq struct {
	Name    string `json:"name" validate:"required,min=1,max=10"`
	Role_Id uint   `json:"role_id" validate:"required"`
}
