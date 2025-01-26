package req

type AgentReq struct {
	Name string `json:"name" validate:"required,min=1,max=10"`
	Role string `json:"role" validate:"required,min=3,max=10"`
}
