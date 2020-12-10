package server

type CreateUserReq struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
