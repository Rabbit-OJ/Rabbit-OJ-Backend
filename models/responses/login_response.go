package responses

type LoginResponse struct {
	Token    string `json:"token"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}
