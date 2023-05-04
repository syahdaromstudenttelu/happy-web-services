package web

type LoginWebResponse struct {
	AuthStatus    bool   `json:"authStatus"`
	AdminUsername string `json:"adminUsername"`
	AdminPassword string `json:"adminPassword"`
}
