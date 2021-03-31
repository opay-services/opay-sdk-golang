package user

type InfoUserReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UserInfo struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

type InfoUserResp struct {
	Code    string   `json:"code"`
	Data    UserInfo `json:"data"`
	Message string   `json:"message"`
}

type UserCreateReq struct {
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	Otp         string `json:"otp"`
}

type SendOptReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

type SendOptResp struct {
	PhoneNumber string `json:"phoneNumber"`
	Status      string `json:"status"`
}

type UserUpdateReq struct {
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
