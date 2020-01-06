package responses

type LoginResponse struct {
	Token    string `json:"token"`
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}
